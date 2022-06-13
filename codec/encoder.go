package codec

import (
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

// Writes list of byte to byte slice
func (e *ByteEncoder) WriteByteList(list [][]byte) {
	var lsize uint32 = uint32(len(list))

	// Writes size of byte list
	e.WriteUint32(lsize)

	for _, item := range list {
		e.WriteBytes(item)
	}
}
