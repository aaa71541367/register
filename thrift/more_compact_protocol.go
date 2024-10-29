package thrift

import (
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

type tMoreCompactType byte

var (
	ttypeToMoreCompactType map[int]tMoreCompactType
	invalidDataLength      error = fmt.Errorf("invalid data length")
)

func init() {
	ttypeToMoreCompactType = map[int]tMoreCompactType{
		0:  0x00,
		2:  0x01,
		3:  0x03,
		6:  0x04,
		8:  0x05,
		10: 0x06,
		4:  0x07,
		11: 0x08,
		15: 0x09,
		14: 0x0A,
		13: 0x0B,
		12: 0x0C,
	}
}

type TMoreCompactProtocolFactory struct {
	cfg *TConfiguration
}

func NewTMoreCompactProtocolFactory() *TMoreCompactProtocolFactory {
	return &TMoreCompactProtocolFactory{&TConfiguration{
		noPropagation: true,
	}}
}

func (p *TMoreCompactProtocolFactory) GetProtocol(trans TTransport) TProtocol {
	return NewTMoreCompactProtocol(trans)
}

type TMoreCompactProtocol struct {
	trans         TRichTransport
	origTransport TTransport

	// Used to keep track of the last field for the current and previous structs,
	// so we can do the delta stuff.
	lastField   []int
	lastFieldId int

	// If we encounter a boolean field begin, save the TField here so it can
	// have the value incorporated.
	booleanFieldName    string
	booleanFieldId      int16
	booleanFieldPending bool

	// If we read a field header, and it's a boolean field, save the boolean
	// value here so that readBool can use it.
	boolValue          bool
	boolValueIsNotNull bool
	buffer             [64]byte

	//tmoreCompact
	f16501s [512]byte
	f16503u []byte
	f16505c []int64
	f16507e int64
	f16508f byte
	f16510h []byte
	f16511i int
	f16512j []string
	f16513k int64
}

// Create a TMoreCompactProtocol given a TTransport
func NewTMoreCompactProtocol(trans TTransport) *TMoreCompactProtocol {
	p := &TMoreCompactProtocol{origTransport: trans, lastField: []int{}}
	if et, ok := trans.(TRichTransport); ok {
		p.trans = et
	} else {
		p.trans = NewTRichTransport(trans)
	}
	p.f16503u = []byte{48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 97, 98, 99, 100, 101, 102}
	p.f16507e = 0
	p.f16501s[26] = 2
	p.f16501s[424] = 3
	p.f16501s[423] = 4
	p.f16501s[214] = 6
	p.f16501s[4] = 8
	p.f16501s[3] = 10
	p.f16501s[51] = 11
	p.f16501s[28] = 12
	p.f16501s[213] = 13
	p.f16501s[212] = 14
	p.f16501s[27] = 15
	p.f16501s[14] = 16
	p.f16501s[11] = 17
	return p

}

//
// Public Writing methods.
//

// Write a message header to the wire. Compact Protocol messages contain the
// protocol version so we can migrate forwards in the future if need be.
func (p *TMoreCompactProtocol) WriteMessageBegin(ctx context.Context, name string, typeId TMessageType, seqid int32) error {
	err := p.writeByteDirect(ctx, COMPACT_PROTOCOL_ID)
	if err != nil {
		return NewTProtocolException(err)
	}
	err = p.writeByteDirect(ctx, (COMPACT_VERSION&COMPACT_VERSION_MASK)|((byte(typeId)<<COMPACT_TYPE_SHIFT_AMOUNT)&COMPACT_TYPE_MASK))
	if err != nil {
		return NewTProtocolException(err)
	}
	_, err = p.writeVarint32(ctx, seqid)
	if err != nil {
		return NewTProtocolException(err)
	}
	e := p.WriteString(ctx, name)
	return e

}

func (p *TMoreCompactProtocol) WriteMessageEnd(ctx context.Context) error { return nil }

// Write a struct begin. This doesn't actually put anything on the wire. We
// use it as an opportunity to put special placeholder markers on the field
// stack so we can get the field id deltas correct.
func (p *TMoreCompactProtocol) WriteStructBegin(ctx context.Context, name string) error {
	p.lastField = append(p.lastField, p.lastFieldId)
	p.lastFieldId = 0
	return nil
}

// Write a struct end. This doesn't actually put anything on the wire. We use
// this as an opportunity to pop the last field from the current struct off
// of the field stack.
func (p *TMoreCompactProtocol) WriteStructEnd(ctx context.Context) error {
	p.lastFieldId = p.lastField[len(p.lastField)-1]
	p.lastField = p.lastField[:len(p.lastField)-1]
	return nil
}

func (p *TMoreCompactProtocol) WriteFieldBegin(ctx context.Context, name string, typeId TType, id int16) error {
	if typeId == BOOL {
		// we want to possibly include the value, so we'll wait.
		p.booleanFieldName, p.booleanFieldId, p.booleanFieldPending = name, id, true
		return nil
	}
	_, err := p.writeFieldBeginInternal(ctx, name, typeId, id, 0xFF)
	return NewTProtocolException(err)
}

// The workhorse of writeFieldBegin. It has the option of doing a
// 'type override' of the type header. This is used specifically in the
// boolean field case.
func (p *TMoreCompactProtocol) writeFieldBeginInternal(ctx context.Context, name string, typeId TType, id int16, typeOverride byte) (int, error) {
	// short lastField = lastField_.pop();

	// if there's a type override, use that.
	var typeToWrite byte
	if typeOverride == 0xFF {
		typeToWrite = byte(p.getCompactType(typeId))
	} else {
		typeToWrite = typeOverride
	}
	// check if we can use delta encoding for the field id
	fieldId := int(id)
	written := 0
	if fieldId > p.lastFieldId && fieldId-p.lastFieldId <= 15 {
		// write them together
		err := p.writeByteDirect(ctx, byte((fieldId-p.lastFieldId)<<4)|typeToWrite)
		if err != nil {
			return 0, err
		}
	} else {
		// write them separate
		err := p.writeByteDirect(ctx, typeToWrite)
		if err != nil {
			return 0, err
		}
		err = p.WriteI16(ctx, id)
		written = 1 + 2
		if err != nil {
			return 0, err
		}
	}

	p.lastFieldId = fieldId
	// p.lastField.Push(field.id);
	return written, nil
}

func (p *TMoreCompactProtocol) WriteFieldEnd(ctx context.Context) error { return nil }

func (p *TMoreCompactProtocol) WriteFieldStop(ctx context.Context) error {
	err := p.writeByteDirect(ctx, STOP)
	return NewTProtocolException(err)
}

func (p *TMoreCompactProtocol) WriteMapBegin(ctx context.Context, keyType TType, valueType TType, size int) error {
	if size == 0 {
		err := p.writeByteDirect(ctx, 0)
		return NewTProtocolException(err)
	}
	_, err := p.writeVarint32(ctx, int32(size))
	if err != nil {
		return NewTProtocolException(err)
	}
	err = p.writeByteDirect(ctx, byte(p.getCompactType(keyType))<<4|byte(p.getCompactType(valueType)))
	return NewTProtocolException(err)
}

func (p *TMoreCompactProtocol) WriteMapEnd(ctx context.Context) error { return nil }

// Write a list header.
func (p *TMoreCompactProtocol) WriteListBegin(ctx context.Context, elemType TType, size int) error {
	_, err := p.writeCollectionBegin(ctx, elemType, size)
	return NewTProtocolException(err)
}

func (p *TMoreCompactProtocol) WriteListEnd(ctx context.Context) error { return nil }

// Write a set header.
func (p *TMoreCompactProtocol) WriteSetBegin(ctx context.Context, elemType TType, size int) error {
	_, err := p.writeCollectionBegin(ctx, elemType, size)
	return NewTProtocolException(err)
}

func (p *TMoreCompactProtocol) WriteSetEnd(ctx context.Context) error { return nil }

func (p *TMoreCompactProtocol) WriteBool(ctx context.Context, value bool) error {
	v := byte(COMPACT_BOOLEAN_FALSE)
	if value {
		v = byte(COMPACT_BOOLEAN_TRUE)
	}
	if p.booleanFieldPending {
		// we haven't written the field header yet
		_, err := p.writeFieldBeginInternal(ctx, p.booleanFieldName, BOOL, p.booleanFieldId, v)
		p.booleanFieldPending = false
		return NewTProtocolException(err)
	}
	// we're not part of a field, so just write the value.
	err := p.writeByteDirect(ctx, v)
	return NewTProtocolException(err)
}

// Write a byte. Nothing to see here!
func (p *TMoreCompactProtocol) WriteByte(ctx context.Context, value int8) error {
	err := p.writeByteDirect(ctx, byte(value))
	return NewTProtocolException(err)
}

// Write an I16 as a zigzag varint.
func (p *TMoreCompactProtocol) WriteI16(ctx context.Context, value int16) error {
	_, err := p.writeVarint32(ctx, p.int32ToZigzag(int32(value)))
	return NewTProtocolException(err)
}

// Write an i32 as a zigzag varint.
func (p *TMoreCompactProtocol) WriteI32(ctx context.Context, value int32) error {
	_, err := p.writeVarint32(ctx, p.int32ToZigzag(value))
	return NewTProtocolException(err)
}

// Write an i64 as a zigzag varint.
func (p *TMoreCompactProtocol) WriteI64(ctx context.Context, value int64) error {
	_, err := p.writeVarint64(ctx, p.int64ToZigzag(value))
	return NewTProtocolException(err)
}

// Write a double to the wire as 8 bytes.
func (p *TMoreCompactProtocol) WriteDouble(ctx context.Context, value float64) error {
	buf := p.buffer[0:8]
	binary.LittleEndian.PutUint64(buf, math.Float64bits(value))
	_, err := p.trans.Write(buf)
	return NewTProtocolException(err)
}

// Write a string to the wire with a varint size preceding.
func (p *TMoreCompactProtocol) WriteString(ctx context.Context, value string) error {
	_, e := p.writeVarint32(ctx, int32(len(value)))
	if e != nil {
		return NewTProtocolException(e)
	}
	if len(value) > 0 {
	}
	_, e = p.trans.WriteString(value)
	return e
}

// Write a byte array, using a varint for the size.
func (p *TMoreCompactProtocol) WriteBinary(ctx context.Context, bin []byte) error {
	_, e := p.writeVarint32(ctx, int32(len(bin)))
	if e != nil {
		return NewTProtocolException(e)
	}
	if len(bin) > 0 {
		_, e = p.trans.Write(bin)
		return NewTProtocolException(e)
	}
	return nil
}

//
// Reading methods.
//

func (p *TMoreCompactProtocol) readBinary(ctx context.Context, i int) ([]byte, error) {
	data := []byte{}
	for s := 0; s < i; s++ {
		b, err := p.readByteDirect(ctx)
		if err != nil {
			return data, err
		}
		data = append(data, b)
	}
	return data, nil
}

// Read a message header.
func (p *TMoreCompactProtocol) ReadMessageBegin(ctx context.Context) (name string, typeId TMessageType, seqId int32, err error) {
	d, err := p.readByteDirect(ctx)
	if err != nil {
		return name, typeId, seqId, err
	}
	if d == 0x92 {
		d2, err := p.readByteDirect(ctx)
		if err != nil {
			return name, typeId, seqId, err
		}
		i := 0
		b := d2 & 31
		if b == 1 {
			b2 := (d2 >> 5) & 3
			C, err := p.readVarint32(ctx)
			if err != nil {
				return name, typeId, seqId, err
			}
			p.f16513k = 0
			p.f16511i = 0
			C2, err := p.readVarint32(ctx)
			if err != nil {
				return name, typeId, seqId, err
			}
			p.f16510h = make([]byte, C2*2)
			i3 := 0
			i4 := 0
			for m := 0; m < int(C2); m++ {
				b4, err := p.readByteDirect(ctx)
				if err != nil {
					return name, typeId, seqId, err
				}
				i5 := 128
				for m := 0; m < 8; m++ {
					if b4&byte(i5) == 0 {
						i3 = i3*2 + i
					} else {
						i3 = i3*2 + 2
					}
					if p.f16501s[i3] != 0 {
						if i4 >= len(p.f16510h) {
							p.f16510h = append(p.f16510h, 0, 0, 0, 0)
						}
						p.f16510h[i4] = p.f16501s[i3]
						i4 += 1
						i3 = 0
					}
					i5 >>= 1
					i = 1
				}
				i = 1
			}
			C3, err := p.readVarint32(ctx)
			if err != nil {
				return name, typeId, seqId, err
			}
			p.f16512j = make([]string, C3)
			for m := 0; m < int(C3); m++ {
				b5, err := p.readBinary(ctx, 17)
				if err != nil {
					return name, typeId, seqId, err
				}
				bArr3 := make([]byte, 33)
				bArr3[0] = b5[0]
				for i10 := 1; i10 < len(b5); i10++ {
					i11 := i10 * 2
					i12 := i11 - 1
					i13 := (b5[i10] >> 4) & 15
					bArr3[i12] = p.f16503u[i13]
					bArr3[i11] = p.f16503u[b5[i10]&15]
				}
				p.f16512j[m] = string(bArr3)
			}
			//fmt.Println(TMessageType(b2), C)
			return "", TMessageType(b2), C, nil
		}
		return "", typeId, seqId, NewTProtocolExceptionWithType(BAD_VERSION, fmt.Errorf("Expected version %02x but got %02x", COMPACT_VERSION, b))
	}
	return "", typeId, seqId, NewTProtocolExceptionWithType(BAD_VERSION, fmt.Errorf("Expected protocol id %02x but got %02x", 0x92, d))
}

func (p *TMoreCompactProtocol) ReadMessageEnd(ctx context.Context) error { return nil }

// Read a struct begin. There's nothing on the wire for this, but it is our
// opportunity to push a new struct begin marker onto the field stack.
func (p *TMoreCompactProtocol) ReadStructBegin(ctx context.Context) (name string, err error) {
	p.lastField = append(p.lastField, p.lastFieldId)
	p.lastFieldId = 0
	p.f16505c = append(p.f16505c, p.f16507e)
	p.f16507e, err = p.readVarint64(ctx)
	return "", err
}

// Doesn't actually consume any wire data, just removes the last field for
// this struct from the field stack.
func (p *TMoreCompactProtocol) ReadStructEnd(ctx context.Context) error {
	if len(p.lastField) != 0 {
		size := len(p.lastField) - 1
		p.lastFieldId = p.lastField[size]
		p.lastField = p.lastField[:size]
	} else {
		p.lastFieldId = 0
	}
	if len(p.f16505c) != 0 {
		size := len(p.f16505c) - 1
		p.f16507e = p.f16505c[size]
		p.f16505c = p.f16505c[:size]
	} else {
		p.f16507e = 0
	}
	return nil
}

// Read a field header off the wire.
func (p *TMoreCompactProtocol) ReadFieldBegin(ctx context.Context) (name string, typeId TType, id int16, err error) {
	s := p.lastFieldId
	if s > 63 || (uint64(p.f16507e)>>s) == 0 {
		return
	}
	for {
		if s > 63 {
			break
		} else if p.f16507e&(1<<s) != 0 {
			p.lastFieldId = s
			break
		} else {
			s += 1
		}
	}
	b := p.f16510h[p.f16511i]
	p.f16511i += 1
	p.f16508f = b
	if b == 16 || b == 17 {
		b = 11
	}
	s2 := p.lastFieldId
	p.lastFieldId += 1
	//fmt.Println(name, TType(b), int16(s2))
	return name, TType(b), int16(s2), nil
}

func (p *TMoreCompactProtocol) ReadFieldEnd(ctx context.Context) error {
	return nil
}

// Read a map header off the wire. If the size is zero, skip reading the key
// and value type. This means that 0-length maps will yield TMaps without the
// "correct" types.
func (p *TMoreCompactProtocol) ReadMapBegin(ctx context.Context) (keyType TType, valueType TType, size int, err error) {
	size32, e := p.readVarint32(ctx)
	if e != nil {
		err = NewTProtocolException(e)
		return
	}
	if size32 < 0 {
		err = invalidDataLength
		return
	}
	size = int(size32)

	keyAndValueType := byte(STOP)
	if size != 0 {
		keyAndValueType, err = p.readByteDirect(ctx)
		if err != nil {
			return
		}
	}
	keyType, _ = p.getTType(tCompactType(keyAndValueType >> 4))
	valueType, _ = p.getTType(tCompactType(keyAndValueType & 0xf))
	return
}

func (p *TMoreCompactProtocol) ReadMapEnd(ctx context.Context) error { return nil }

// Read a list header off the wire. If the list size is 0-14, the size will
// be packed into the element type header. If it's a longer list, the 4 MSB
// of the element type header will be 0xF, and a varint will follow with the
// true size.
func (p *TMoreCompactProtocol) ReadListBegin(ctx context.Context) (elemType TType, size int, err error) {
	size_and_type, err := p.readByteDirect(ctx)
	if err != nil {
		return
	}
	size = int((size_and_type >> 4) & 0x0f)
	if size == 15 {
		size2, e := p.readVarint32(ctx)
		if e != nil {
			err = NewTProtocolException(e)
			return
		}
		if size2 < 0 {
			err = invalidDataLength
			return
		}
		size = int(size2)
	}
	elemType, e := p.getTType(tCompactType(size_and_type))
	if e != nil {
		err = NewTProtocolException(e)
		return
	}
	return
}

func (p *TMoreCompactProtocol) ReadListEnd(ctx context.Context) error { return nil }

// Read a set header off the wire. If the set size is 0-14, the size will
// be packed into the element type header. If it's a longer set, the 4 MSB
// of the element type header will be 0xF, and a varint will follow with the
// true size.
func (p *TMoreCompactProtocol) ReadSetBegin(ctx context.Context) (elemType TType, size int, err error) {
	return p.ReadListBegin(ctx)
}

func (p *TMoreCompactProtocol) ReadSetEnd(ctx context.Context) error { return nil }

// Read a boolean off the wire. If this is a boolean field, the value should
// already have been read during readFieldBegin, so we'll just consume the
// pre-stored value. Otherwise, read a byte.
func (p *TMoreCompactProtocol) ReadBool(ctx context.Context) (value bool, err error) {
	v, err := p.readByteDirect(ctx)
	return v == COMPACT_BOOLEAN_TRUE, err
}

// Read a single byte off the wire. Nothing interesting here.
func (p *TMoreCompactProtocol) ReadByte(ctx context.Context) (int8, error) {
	v, err := p.readByteDirect(ctx)
	if err != nil {
		return 0, NewTProtocolException(err)
	}
	return int8(v), err
}

// Read an i16 from the wire as a zigzag varint.
func (p *TMoreCompactProtocol) ReadI16(ctx context.Context) (value int16, err error) {
	v, err := p.ReadI32(ctx)
	return int16(v), err
}

// Read an i32 from the wire as a zigzag varint.
func (p *TMoreCompactProtocol) ReadI32(ctx context.Context) (value int32, err error) {
	v, e := p.readVarint32(ctx)
	if e != nil {
		return 0, NewTProtocolException(e)
	}
	value = p.zigzagToInt32(v)
	return value, nil
}

// Read an i64 from the wire as a zigzag varint.
func (p *TMoreCompactProtocol) ReadI64(ctx context.Context) (value int64, err error) {
	v, e := p.readVarint64(ctx)
	if e != nil {
		return 0, NewTProtocolException(e)
	}
	value = p.zigzagToInt64(v)
	return value, nil
}

// No magic here - just read a double off the wire.
func (p *TMoreCompactProtocol) ReadDouble(ctx context.Context) (value float64, err error) {
	longBits := make([]byte, 8)
	_, e := io.ReadFull(p.trans, longBits)
	if e != nil {
		return 0.0, NewTProtocolException(e)
	}
	return math.Float64frombits(p.bytesToUint64(longBits)), nil
}

// Reads a []byte (via readBinary), and then UTF-8 decodes it.
func (p *TMoreCompactProtocol) ReadString(ctx context.Context) (value string, err error) {
	b := p.f16508f
	if b == 16 {
		j, err := p.ReadI64(ctx)
		p.f16513k += j
		return fmt.Sprintf("%d", p.f16513k), err
	} else if b == 17 {
		j, err := p.readVarint32(ctx)
		return p.f16512j[j], err
	} else {
		C, err := p.readVarint32(ctx)
		if C == 0 {
			return "", err
		} else {
			if C > -1 {
				res, err := p.readBinary(ctx, int(C))
				return string(res), err
			} else {
				return "", nil
			}
		}
	}
}

// Read a []byte from the wire.
func (p *TMoreCompactProtocol) ReadBinary(ctx context.Context) (value []byte, err error) {
	length, e := p.readVarint32(ctx)
	if e != nil {
		return nil, NewTProtocolException(e)
	}
	if length == 0 {
		return []byte{}, nil
	}
	if length < 0 {
		return nil, invalidDataLength
	}

	buf := make([]byte, length)
	_, e = io.ReadFull(p.trans, buf)
	return buf, NewTProtocolException(e)
}

func (p *TMoreCompactProtocol) Flush(ctx context.Context) (err error) {
	return NewTProtocolException(p.trans.Flush(ctx))
}

func (p *TMoreCompactProtocol) Skip(ctx context.Context, fieldType TType) (err error) {
	return SkipDefaultDepth(ctx, p, fieldType)
}

func (p *TMoreCompactProtocol) Transport() TTransport {
	return p.origTransport
}

//
// Internal writing methods
//

// Abstract method for writing the start of lists and sets. List and sets on
// the wire differ only by the type indicator.
func (p *TMoreCompactProtocol) writeCollectionBegin(ctx context.Context, elemType TType, size int) (int, error) {
	if size <= 14 {
		return 1, p.writeByteDirect(ctx, byte(int32(size<<4)|int32(p.getCompactType(elemType))))
	}
	err := p.writeByteDirect(ctx, 0xf0|byte(p.getCompactType(elemType)))
	if err != nil {
		return 0, err
	}
	m, err := p.writeVarint32(ctx, int32(size))
	return 1 + m, err
}

// Write an i32 as a varint. Results in 1-5 bytes on the wire.
// TODO(pomack): make a permanent buffer like writeVarint64?
func (p *TMoreCompactProtocol) writeVarint32(ctx context.Context, n int32) (int, error) {
	i32buf := p.buffer[0:5]
	idx := 0
	for {
		if (n & ^0x7F) == 0 {
			i32buf[idx] = byte(n)
			idx++
			// p.writeByteDirect(byte(n));
			break
			// return;
		} else {
			i32buf[idx] = byte((n & 0x7F) | 0x80)
			idx++
			// p.writeByteDirect(byte(((n & 0x7F) | 0x80)));
			u := uint32(n)
			n = int32(u >> 7)
		}
	}
	return p.trans.Write(i32buf[0:idx])
}

// Write an i64 as a varint. Results in 1-10 bytes on the wire.
func (p *TMoreCompactProtocol) writeVarint64(ctx context.Context, n int64) (int, error) {
	varint64out := p.buffer[0:10]
	idx := 0
	for {
		if (n & ^0x7F) == 0 {
			varint64out[idx] = byte(n)
			idx++
			break
		} else {
			varint64out[idx] = byte((n & 0x7F) | 0x80)
			idx++
			u := uint64(n)
			n = int64(u >> 7)
		}
	}
	return p.trans.Write(varint64out[0:idx])
}

// Convert l into a zigzag long. This allows negative numbers to be
// represented compactly as a varint.
func (p *TMoreCompactProtocol) int64ToZigzag(l int64) int64 {
	return (l << 1) ^ (l >> 63)
}

// Convert l into a zigzag long. This allows negative numbers to be
// represented compactly as a varint.
func (p *TMoreCompactProtocol) int32ToZigzag(n int32) int32 {
	return (n << 1) ^ (n >> 31)
}

func (p *TMoreCompactProtocol) fixedUint64ToBytes(n uint64, buf []byte) {
	binary.LittleEndian.PutUint64(buf, n)
}

func (p *TMoreCompactProtocol) fixedInt64ToBytes(n int64, buf []byte) {
	binary.LittleEndian.PutUint64(buf, uint64(n))
}

// Writes a byte without any possibility of all that field header nonsense.
// Used internally by other writing methods that know they need to write a byte.
func (p *TMoreCompactProtocol) writeByteDirect(ctx context.Context, b byte) error {
	return p.trans.WriteByte(b)
}

// Writes a byte without any possibility of all that field header nonsense.
func (p *TMoreCompactProtocol) writeIntAsByteDirect(ctx context.Context, n int) (int, error) {
	return 1, p.writeByteDirect(ctx, byte(n))
}

//
// Internal reading methods
//

// Read an i32 from the wire as a varint. The MSB of each byte is set
// if there is another byte to follow. This can read up to 5 bytes.
func (p *TMoreCompactProtocol) readVarint32(ctx context.Context) (int32, error) {
	// if the wire contains the right stuff, this will just truncate the i64 we
	// read and get us the right sign.
	v, err := p.readVarint64(ctx)
	return int32(v), err
}

// Read an i64 from the wire as a proper varint. The MSB of each byte is set
// if there is another byte to follow. This can read up to 10 bytes.
func (p *TMoreCompactProtocol) readVarint64(ctx context.Context) (int64, error) {
	shift := uint(0)
	result := int64(0)
	for {
		b, err := p.readByteDirect(ctx)
		if err != nil {
			return 0, err
		}
		result |= int64(b&0x7f) << shift
		if (b & 0x80) != 0x80 {
			break
		}
		shift += 7
	}
	return result, nil
}

// Read a byte, unlike ReadByte that reads Thrift-byte that is i8.
func (p *TMoreCompactProtocol) readByteDirect(ctx context.Context) (byte, error) {
	return p.trans.ReadByte()
}

//
// encoding helpers
//

// Convert from zigzag int to int.
func (p *TMoreCompactProtocol) zigzagToInt32(n int32) int32 {
	u := uint32(n)
	return int32(u>>1) ^ -(n & 1)
}

// Convert from zigzag long to long.
func (p *TMoreCompactProtocol) zigzagToInt64(n int64) int64 {
	u := uint64(n)
	return int64(u>>1) ^ -(n & 1)
}

// Note that it's important that the mask bytes are long literals,
// otherwise they'll default to ints, and when you shift an int left 56 bits,
// you just get a messed up int.
func (p *TMoreCompactProtocol) bytesToInt64(b []byte) int64 {
	return int64(binary.LittleEndian.Uint64(b))
}

// Note that it's important that the mask bytes are long literals,
// otherwise they'll default to ints, and when you shift an int left 56 bits,
// you just get a messed up int.
func (p *TMoreCompactProtocol) bytesToUint64(b []byte) uint64 {
	return binary.LittleEndian.Uint64(b)
}

//
// type testing and converting
//

func (p *TMoreCompactProtocol) isBoolType(b byte) bool {
	return (b&0x0f) == COMPACT_BOOLEAN_TRUE || (b&0x0f) == COMPACT_BOOLEAN_FALSE
}

// Given a tCompactType constant, convert it to its corresponding
// TType value.
func (p *TMoreCompactProtocol) getTType(t tCompactType) (TType, error) {
	switch byte(t) & 0x0f {
	case STOP:
		return STOP, nil
	case COMPACT_BOOLEAN_FALSE, COMPACT_BOOLEAN_TRUE:
		return BOOL, nil
	case COMPACT_BYTE:
		return BYTE, nil
	case COMPACT_I16:
		return I16, nil
	case COMPACT_I32:
		return I32, nil
	case COMPACT_I64:
		return I64, nil
	case COMPACT_DOUBLE:
		return DOUBLE, nil
	case COMPACT_BINARY:
		return STRING, nil
	case COMPACT_LIST:
		return LIST, nil
	case COMPACT_SET:
		return SET, nil
	case COMPACT_MAP:
		return MAP, nil
	case COMPACT_STRUCT:
		return STRUCT, nil
	}
	return STOP, NewTProtocolException(fmt.Errorf("don't know what type: %v", t&0x0f))
}

// Given a TType value, find the appropriate TMoreCompactProtocol.Types constant.
func (p *TMoreCompactProtocol) getCompactType(t TType) tMoreCompactType {
	return ttypeToMoreCompactType[int(t)]
}
