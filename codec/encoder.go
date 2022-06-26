package codec

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type ByteEncoder struct {
	Error   error
	maxSize int
	Bytes   []byte
	offset  int
}

// Returns new ByteEncoder
func NewByteEncoder(size int) *ByteEncoder {
	return &ByteEncoder{
		Error:   nil,
		maxSize: size,
		Bytes:   []byte{},
		offset:  0,
	}
}

func (e *ByteEncoder) expandSize(size int) error {
	var newSize int = size + e.offset

	if newSize > e.maxSize {
		e.Error = errors.New("insufficient space in byte array")
	} else if newSize <= cap(e.Bytes) {
		e.Bytes = e.Bytes[:newSize]
	} else {
		// expand size for byte slice
		e.Bytes = append(e.Bytes[:cap(e.Bytes)], make([]byte, newSize-cap(e.Bytes))...)
	}

	return e.Error
}

// Appends single byte to byte slice
func (e *ByteEncoder) WriteSingleByte(v byte) {
	if e.expandSize(ByteSize) != nil {
		return
	}

	e.Bytes[e.offset] = v
	e.offset += ByteSize
}

// Writes byte of type bool to byte slice
func (e *ByteEncoder) WriteBool(v bool) {
	if e.expandSize(BoolSize) != nil {
		return
	}

	if v {
		e.WriteSingleByte(1)
	} else {
		e.WriteSingleByte(0)
	}
}

// Appends fixed size of bytes to byte slice
func (e *ByteEncoder) WriteBytes(v []byte) {
	var bsize int = len(v)

	if e.expandSize(bsize) != nil {
		return
	}

	copy(e.Bytes[e.offset:], v)
	e.offset += bsize
}

func (e *ByteEncoder) WriteFixedBytes(v []byte, size int) {
	if e.expandSize(size) != nil {
		return
	}

	copy(e.Bytes[e.offset:], v)
	e.offset += size
}

// Writes bytes of type uint16 to byte slice
func (e *ByteEncoder) WriteUint16(v uint16) {
	if e.expandSize(Uint16Size) != nil {
		return
	}

	binary.BigEndian.PutUint16(e.Bytes[e.offset:], v)
	e.offset += Uint16Size
}

// Writes bytes of type uint32 to byte slice
func (e *ByteEncoder) WriteUint32(v uint32) {
	if e.expandSize(Uint32Size) != nil {
		return
	}

	binary.BigEndian.PutUint32(e.Bytes[e.offset:], v)
	e.offset += Uint32Size
}

// Writes bytes of type uint64 to byte slice
func (e *ByteEncoder) WriteUint64(v uint64) {
	if e.expandSize(Uint64Size) != nil {
		return
	}

	binary.BigEndian.PutUint64(e.Bytes[e.offset:], v)
	e.offset += Uint64Size
}

type EncodeBuffer struct {
	err  error
	buff *bytes.Buffer
}

func NewEncodeBuffer() *EncodeBuffer {
	return &EncodeBuffer{
		err:  nil,
		buff: &bytes.Buffer{},
	}
}

func (e *EncodeBuffer) Buffer() *bytes.Buffer {
	return e.buff
}

func (e *EncodeBuffer) Error() error {
	return e.err
}

func (e *EncodeBuffer) WriteUint64(v uint64) {
	if e.err != nil {
		return
	}

	e.err = binary.Write(e.buff, binary.BigEndian, v)
}

func (e *EncodeBuffer) WriteUint32(v uint32) {
	if e.err != nil {
		return
	}

	e.err = binary.Write(e.buff, binary.BigEndian, v)
}

func (e *EncodeBuffer) WriteUint16(v uint16) {
	if e.err != nil {
		return
	}

	e.err = binary.Write(e.buff, binary.BigEndian, v)
}

func (e *EncodeBuffer) WriteSingleByte(v byte) {
	if e.err != nil {
		return
	}

	e.err = binary.Write(e.buff, binary.BigEndian, v)
}

func (e *EncodeBuffer) WriteBool(v bool) {
	if e.err != nil {
		return
	}

	e.err = binary.Write(e.buff, binary.BigEndian, v)
}

func (e *EncodeBuffer) WriteFixedBytes(v []byte) {
	if e.err != nil {
		return
	}

	e.err = binary.Write(e.buff, binary.BigEndian, v)
}

func (e *EncodeBuffer) Reset() {
	e.err = nil
	e.buff.Reset()
}
