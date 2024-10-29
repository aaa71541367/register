/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License. You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package thrift

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/valyala/fasthttp"
	//"strconv"
)

// Default to using the shared http client. Library users are
// free to change this global client or specify one through
// TFastHttpClientOptions.

//var DefaulTFastHttpClient *fasthttp.Client = new(fasthttp.Client)

type httpClient interface {
	Do(req *fasthttp.Request, resp *fasthttp.Response) error
	DoTimeout(req *fasthttp.Request, resp *fasthttp.Response, timeout time.Duration) error
}

type TFastHttpClient struct {
	client        httpClient
	request       *fasthttp.Request
	requestBuffer *bytes.Buffer
	respReader    *bytes.Reader
}

type TFastHttpClientTransportFactory struct {
	options TFastHttpClientOptions
	url     string
}

func (p *TFastHttpClientTransportFactory) GetTransport(trans TTransport) (TTransport, error) {
	if trans != nil {
		return NewTFastHttpClientWithOptions(p.url, p.options)
	}
	return NewTFastHttpClientWithOptions(p.url, p.options)
}

type TFastHttpClientOptions struct {
	// If nil, DefaulTFastHttpClient is used
	Client httpClient
}

func NewTFastHttpClientTransportFactory(url string) *TFastHttpClientTransportFactory {
	return NewTFastHttpClientTransportFactoryWithOptions(url, TFastHttpClientOptions{})
}

func NewTFastHttpClientTransportFactoryWithOptions(url string, options TFastHttpClientOptions) *TFastHttpClientTransportFactory {
	return &TFastHttpClientTransportFactory{url: url, options: options}
}

func NewTFastHttpClientWithOptions(urlstr string, options TFastHttpClientOptions) (TTransport, error) {
	client := options.Client
	if client == nil {
		client = new(fasthttp.Client)
	}
	request := fasthttp.AcquireRequest()
	request.Header.SetContentType("application/x-thrift")
	request.Header.SetMethod("POST")
	request.SetRequestURI(urlstr)
	return &TFastHttpClient{client: client, request: request, requestBuffer: new(bytes.Buffer), respReader: new(bytes.Reader)}, nil
}

func NewTFastHttpClient(urlstr string) (TTransport, error) {
	return NewTFastHttpClientWithOptions(urlstr, TFastHttpClientOptions{})
}

// Set the HTTP Header for this specific Thrift Transport
// It is important that you first assert the TTransport as a TFastHttpClient type
// like so:
//
// httpTrans := trans.(TFastHttpClient)
// httpTrans.SetHeader("User-Agent","Thrift Client 1.0")
func (p *TFastHttpClient) SetHeader(key string, value string) {
	if key != "" && value != "" {
		p.request.Header.Set(key, value)
	}
}

// Get the HTTP Header represented by the supplied Header Key for this specific Thrift Transport
// It is important that you first assert the TTransport as a TFastHttpClient type
// like so:
//
// httpTrans := trans.(TFastHttpClient)
// hdrValue := httpTrans.GetHeader("User-Agent")
func (p *TFastHttpClient) GetHeader(key string) string {
	return string(p.request.Header.Peek(key))
}

// Deletes the HTTP Header given a Header Key for this specific Thrift Transport
// It is important that you first assert the TTransport as a TFastHttpClient type
// like so:
//
// httpTrans := trans.(TFastHttpClient)
// httpTrans.DelHeader("User-Agent")
func (p *TFastHttpClient) DelHeader(key string) {
	p.request.Header.Del(key)
}

func (p *TFastHttpClient) Open() error {
	// do nothing
	return nil
}

func (p *TFastHttpClient) IsOpen() bool {
	return p.requestBuffer != nil
}

func (p *TFastHttpClient) closeResponse() error {
	if p.respReader != nil {
		io.Copy(ioutil.Discard, p.respReader)
	}
	return nil
}

func (p *TFastHttpClient) Close() error {
	if p.requestBuffer != nil {
		p.requestBuffer.Reset()
	}
	return p.closeResponse()
}

func (p *TFastHttpClient) Read(buf []byte) (int, error) {
	n, err := p.respReader.Read(buf)
	if n > 0 && (err == nil || err == io.EOF) {
		return n, nil
	}
	return n, NewTTransportExceptionFromError(err)
}

func (p *TFastHttpClient) ReadByte() (c byte, err error) {
	return readByte(p.respReader)
}

func (p *TFastHttpClient) Write(buf []byte) (int, error) {
	return p.requestBuffer.Write(buf)
}

func (p *TFastHttpClient) WriteByte(c byte) error {
	return p.requestBuffer.WriteByte(c)
}

func (p *TFastHttpClient) WriteString(s string) (n int, err error) {
	return p.requestBuffer.WriteString(s)
}

func (p *TFastHttpClient) SetResponseBody(buf []byte) {
	p.respReader.Reset(buf)
}

func (p *TFastHttpClient) Flush(ctx context.Context) error {
	response := fasthttp.AcquireResponse()
	defer func() {
		// fasthttp.ReleaseRequest(p.request)
		fasthttp.ReleaseResponse(response)
	}()
	p.request.SetBody(p.requestBuffer.Bytes())
	// fmt.Println(p.requestBuffer.Bytes())
	// fmt.Println(p.request.String())
	err := p.client.Do(p.request, response)
	// fmt.Println(p.response.Body())
	// fmt.Println(response)
	p.respReader.Reset(response.Body())
	p.requestBuffer.Reset()

	if err != nil {
		fmt.Println("flush", err)
		return NewTTransportExceptionFromError(err)
	} else if response.StatusCode() != 200 {
		// Close the response to avoid leaking file descriptors. closeResponse does
		// more than just call Close(), so temporarily assign it and reuse the logic.
		return NewTTransportException(UNKNOWN_TRANSPORT_EXCEPTION, "HTTP Response code: "+strconv.Itoa(response.StatusCode()))
	}
	return nil
}

func (p *TFastHttpClient) RemainingBytes() (num_bytes uint64) {
	len := p.respReader.Size()
	if len >= 0 {
		return uint64(len)
	}

	const maxSize = ^uint64(27)
	return maxSize // the thruth is, we just don't know unless framed is used
}
