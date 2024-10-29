package LineThrift

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"

	"github.com/shillbie/register/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

type CompactMessageService interface {
	// Parameters:
	//  - msgTo
	//  - text
	//  - toType
	SendCompactMessage(ctx context.Context, reqSeq int32, msgTo string, text string, toType MIDType) (message *Message, err error)
	// Parameters:
	//  - msgTo
	//  - toType
	//  - e2eeVersion
	//  - salt
	//  - ciphertext
	//  - nonce
	//  - fromKeyId
	//  - toKeyId
	SendEncryptCompactMessage(ctx context.Context, msgTo string, toType MIDType, e2eeVersion int8, salt, ciphertext, nonce []byte, fromKeyId, toKeyId int32) (message *Message, err error)
}

type CompactMessageServiceClient struct {
	iprot, oprot thrift.TProtocol
}

func NewCompactMessageServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *CompactMessageServiceClient {
	iprot := f.GetProtocol(t)
	oprot := f.GetProtocol(t)
	return &CompactMessageServiceClient{
		iprot: iprot,
		oprot: oprot,
	}
}

func NewCompactMessageServiceClientProtocol(iprot thrift.TProtocol, oprot thrift.TProtocol) *CompactMessageServiceClient {
	return &CompactMessageServiceClient{
		iprot: iprot,
		oprot: oprot,
	}
}

// Parameters:
//  - msgTo
//  - text
//  - toType
func (p *CompactMessageServiceClient) SendCompactText(ctx context.Context, msgTo string, text string, toType MIDType) (message *Message, err error) {
	if len(msgTo) != 33 {
		return nil, thrift.PrependError(fmt.Sprintf("%T write message error: ", p), fmt.Errorf("invalid target length"))
	}
	message = &Message{
		To:     msgTo,
		Text:   text,
		ToType: toType,
	}
	if err = p.oprot.WriteByte(ctx, 2); err != nil {
		return
	}
	if err = p.oprot.WriteByte(ctx, int8(toType)); err != nil {
		return
	}
	if err = p.oprot.WriteI32(ctx, 0); err != nil {
		return
	}
	unHexMid, _ := hex.DecodeString(msgTo[1:])
	for _, b := range unHexMid {
		if err = p.oprot.WriteByte(ctx, int8(b)); err != nil {
			return
		}
	}
	if err = p.oprot.WriteString(ctx, text); err != nil {
		return
	}
	if err = p.oprot.WriteByte(ctx, 2); err != nil {
		return
	}
	if err = p.oprot.Flush(ctx); err != nil {
		return
	}
	if err = p.read(ctx, message); err != nil {
		return
	}
	return
}

// Parameters:
//  - msgTo
//  - toType
//  - e2eeVersion
//  - salt
//  - ciphertext
//  - nonce
//  - fromKeyId
//  - toKeyId
func (p *CompactMessageServiceClient) SendEncryptCompactMessage(ctx context.Context, msgTo string, toType MIDType, e2eeVersion int8, salt, ciphertext, nonce []byte, fromKeyId, toKeyId int32) (message *Message, err error) {
	if len(msgTo) != 33 {
		return nil, thrift.PrependError(fmt.Sprintf("%T write message error: ", p), fmt.Errorf("invalid target length"))
	}
	message = new(Message)
	message.To = msgTo
	message.ToType = toType

	if err = p.oprot.WriteByte(ctx, 5); err != nil {
		return
	}
	if err = p.oprot.WriteByte(ctx, int8(toType)); err != nil {
		return
	}
	if err = p.oprot.WriteI32(ctx, 0); err != nil {
		return
	}
	unHexMid, _ := hex.DecodeString(msgTo[1:])
	for _, b := range unHexMid {
		if err = p.oprot.WriteByte(ctx, int8(b)); err != nil {
			return
		}
	}
	if err = p.oprot.WriteByte(ctx, e2eeVersion); err != nil {
		return
	}
	if err = p.oprot.WriteBinary(ctx, salt); err != nil {
		return
	}
	if err = p.oprot.WriteBinary(ctx, ciphertext); err != nil {
		return
	}
	if err = p.oprot.WriteBinary(ctx, nonce); err != nil {
		return
	}
	if err = p.oprot.WriteI32(ctx, fromKeyId); err != nil {
		return
	}
	if err = p.oprot.WriteI32(ctx, toKeyId); err != nil {
		return
	}
	if err = p.oprot.Flush(ctx); err != nil {
		return
	}
	if err = p.read(ctx, message); err != nil {
		return
	}
	return message, nil
}

func (p *CompactMessageServiceClient) read(ctx context.Context, message *Message) (err error) {
	if message == nil {
		message = new(Message)
	}
	var sent bool
	if sent, err = p.iprot.ReadBool(ctx); err != nil {
		return
	}
	if !sent {
		var errorCode int32
		if errorCode, err = p.iprot.ReadI32(ctx); err != nil {
			return
		}
		err = &TalkException{Code: ErrorCode(errorCode)}
		return
	}
	if _, err = p.iprot.ReadI32(ctx); err != nil {
		return
	}
	var msgID int64
	if msgID, err = p.iprot.ReadI64(ctx); err != nil {
		return
	}
	if message.CreatedTime, err = p.iprot.ReadI64(ctx); err != nil {
		return
	}
	message.ID = strconv.FormatInt(msgID, 10)
	return
}
