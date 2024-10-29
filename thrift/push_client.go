package thrift

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"net/url"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/hpack"
)

type PushClient struct {
	reqid          int
	client         net.Conn
	framer         *http2.Framer
	requestBuffer  *bytes.Buffer
	response       chan error
	responseBuffer *bytes.Buffer
	dialer         func() (net.Conn, error)
	header         map[string]string
}

type PushClientOptions struct {
	Client *tls.Conn
}

func NewPushClient(urlstr string) (TTransport, error) {
	parsedURL, err := url.Parse(urlstr)
	if err != nil {
		return nil, err
	}
	dialer := func() (net.Conn, error) {
		client, err := tls.Dial("tcp", parsedURL.Host+":443", &tls.Config{
			NextProtos: []string{"h2"},
		})
		if err != nil {
			return nil, err
		}
		err = client.Handshake()
		return client, err
	}
	return &PushClient{dialer: dialer, client: nil, framer: nil, requestBuffer: bytes.NewBuffer([]byte{}), response: make(chan error), responseBuffer: bytes.NewBuffer([]byte{}), header: map[string]string{":method": "POST", ":authority": parsedURL.Host, ":scheme": "https", ":path": parsedURL.Path + "?" + parsedURL.RawQuery}, reqid: 0}, nil
}

func NewPushClientWithDialer(urlstr string, dialer func() (net.Conn, error)) (TTransport, error) {
	trans, err := NewPushClient(urlstr)
	if err == nil {
		trans.(*PushClient).dialer = dialer
	}
	return trans, err
}

func (p *PushClient) buildRequest(service int, data []byte) []byte {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, uint16(len(data)))
	buf = append(buf, byte(service))
	buf = append(buf, data...)
	return buf
}

func (p *PushClient) buildSignOnRequest(data []byte) []byte {
	res := []byte{0, 0, 5, 0, 0, 0}
	binary.BigEndian.PutUint16(res[0:2], uint16(p.reqid))
	binary.BigEndian.PutUint16(res[4:6], uint16(len(data)))
	res = append(res, data...)
	return res
}

func (p *PushClient) SetHeader(key string, value string) {
	p.header[key] = value
}

func (p *PushClient) GetHeader(key string) string {
	return p.header[key]
}

func (p *PushClient) DelHeader(key string) {
	delete(p.header, key)
}

func (p *PushClient) Open() (err error) {
	p.client, err = p.dialer()
	if err != nil {
		return err
	}
	io.WriteString(p.client, http2.ClientPreface)
	framer := http2.NewFramer(p.client, p.client)
	framer.WriteSettings(
		http2.Setting{ID: http2.SettingHeaderTableSize, Val: 4096},
		http2.Setting{ID: http2.SettingEnablePush, Val: 1},
		http2.Setting{ID: http2.SettingMaxConcurrentStreams, Val: 100},
		http2.Setting{ID: http2.SettingInitialWindowSize, Val: 65535},
		http2.Setting{ID: http2.SettingMaxFrameSize, Val: 16384},
		http2.Setting{ID: http2.SettingMaxHeaderListSize, Val: 65536},
		http2.Setting{ID: 8, Val: 0},
	)
	var headers bytes.Buffer
	enc := hpack.NewEncoder(&headers)
	for key, val := range p.header {
		enc.WriteField(hpack.HeaderField{Name: key, Value: val})
	}
	err = framer.WriteHeaders(http2.HeadersFrameParam{
		EndHeaders:    true,
		BlockFragment: headers.Bytes(),
		StreamID:      1,
		EndStream:     false,
	})
	if err != nil {
		return err
	}
	framer.WriteData(1, false, p.buildRequest(0, []byte{0, 0, 30}))
	p.framer = framer
	return nil
}

func (p *PushClient) IsOpen() bool {
	return p.client != nil
}

func (p *PushClient) Close() error {
	p.client.Close()
	p.client = nil
	return nil
}

func (p *PushClient) Read(buf []byte) (int, error) {
	if p.responseBuffer == nil {
		return 0, NewTTransportException(NOT_OPEN, "Response buffer is empty, no request.")
	}
	n, err := p.responseBuffer.Read(buf)
	if n > 0 && (err == nil || errors.Is(err, io.EOF)) {
		return n, nil
	}
	return n, NewTTransportExceptionFromError(err)
}

func (p *PushClient) ReadByte() (c byte, err error) {
	if p.responseBuffer == nil {
		return 0, NewTTransportException(NOT_OPEN, "Response buffer is empty, no request.")
	}
	return readByte(p.responseBuffer)
}

func (p *PushClient) Write(buf []byte) (int, error) {
	if p.requestBuffer == nil {
		return 0, NewTTransportException(NOT_OPEN, "Request buffer is nil, connection may have been closed.")
	}
	return p.requestBuffer.Write(buf)
}

func (p *PushClient) WriteByte(c byte) error {
	if p.requestBuffer == nil {
		return NewTTransportException(NOT_OPEN, "Request buffer is nil, connection may have been closed.")
	}
	return p.requestBuffer.WriteByte(c)
}

func (p *PushClient) WriteString(s string) (n int, err error) {
	if p.requestBuffer == nil {
		return 0, NewTTransportException(NOT_OPEN, "Request buffer is nil, connection may have been closed.")
	}
	return p.requestBuffer.WriteString(s)
}

func (p *PushClient) Flush(ctx context.Context) error {
	if !p.IsOpen() {
		if err := p.Open(); err != nil {
			return err
		}
	}
	p.reqid += 1
	p.framer.WriteData(1, false, p.buildRequest(2, p.buildSignOnRequest(p.requestBuffer.Bytes())))
	io.WriteString(p.client, "")
	p.requestBuffer.Reset()
	go func(ctx context.Context) {
		p.responseBuffer.Reset()
		data := []byte{}
		for {
			f, err := p.framer.ReadFrame()
			if err != nil || ctx.Err() != nil {
				p.response <- err
				return
			}
			switch frame := f.(type) {
			case *http2.DataFrame:
				data = append(data, frame.Data()...)
				if len(data) < 4 {
					p.response <- NewTTransportException(UNKNOWN_TRANSPORT_EXCEPTION, fmt.Sprintf("[CONN] Invalid Packet: %v", data))
					return
				}
				if data[2] == 1 {
					if data[3] == 2 {
						p.framer.WriteData(1, false, append([]byte{0, 3, 1, 1}, data[4:6]...))
						io.WriteString(p.client, "")
						data = data[:0]
					} else {
						p.response <- NewTTransportException(UNKNOWN_TRANSPORT_EXCEPTION, fmt.Sprintf("invalid ping id: %d", data[3]))
						return
					}
				} else if data[2] == 3 {
					if binary.BigEndian.Uint16(data[0:2]) == uint16(len(data[3:])) {
						p.responseBuffer.Write(data[5:])
						p.response <- nil
						return
					}
				} else {
					p.response <- NewTTransportException(UNKNOWN_TRANSPORT_EXCEPTION, fmt.Sprintf("[CONN] Not Impl: %v", data))
					return
				}
			case *http2.RSTStreamFrame:
				p.response <- NewTTransportException(UNKNOWN_TRANSPORT_EXCEPTION, "Server reset connection")
				return
			case *http2.GoAwayFrame:
				p.response <- NewTTransportException(UNKNOWN_TRANSPORT_EXCEPTION, "Server sent goaway")
				return
			case *http2.SettingsFrame:
				if frame.IsAck() {
					p.framer.WriteSettingsAck()
				}
			case *http2.HeadersFrame, *http2.WindowUpdateFrame:
				io.WriteString(p.client, "")
			default:
				io.WriteString(p.client, "")
				p.response <- NewTTransportException(UNKNOWN_TRANSPORT_EXCEPTION, fmt.Sprintf("unknown frame: %v", frame))
				return
			}
		}
	}(ctx)
	var err error
	select {
	case err = <-p.response:
	case <-ctx.Done():
		err = ctx.Err()
	}
	if err != nil {
		p.Close()
	}
	return err
}

func (p *PushClient) RemainingBytes() (num_bytes uint64) {
	len := p.requestBuffer.Len()
	if len >= 0 {
		return uint64(len)
	}

	const maxSize = ^uint64(0)
	return maxSize // the truth is, we just don't know unless framed is used
}
