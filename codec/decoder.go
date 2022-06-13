package codec

import (
	"encoding/binary"
	"errors"
)

type ByteDecoder struct {
	Error  error
	Bytes  []byte
	offset int
}

// Returns new ByteDecoder
func NewByteDecoder(data []byte) *ByteDecoder {
	return &ByteDecoder{
		Error:  nil,
		Bytes:  data,
		offset: 0,
	}
}

func (d *ByteDecoder) checkSize(size int) error {
	switch {
	case size < 0:
		d.Error = errors.New("invalid input size")
	case d.offset < 0:
		d.Error = errors.New("negative offset")
	case len(d.Bytes)-d.offset < size:
		d.Error = errors.New("insufficient input legnth")
	}
	return d.Error
}

// Reads single byte from byte slice
func (d *ByteDecoder) ReadSingleByte(r *byte) {
	if d.checkSize(ByteSize) != nil {
		return
	}

	*r = d.Bytes[d.offset]
	d.offset += ByteSize
}

// Reads fixed size of bytes from byte slice
func (d *ByteDecoder) ReadBytes(r *[]byte, size int) {
	if d.checkSize(size) != nil {
		return
	}

	*r = d.Bytes[d.offset : d.offset+size]
	d.offset += size
}

// Reads byte of bool type from byte slice
func (d *ByteDecoder) ReadBool(r *bool) {
	var b byte

	if d.checkSize(BoolSize) != nil {
		return
	}

	switch d.ReadSingleByte(&b); {
	case b == 1:
		*r = true
	case b == 0:
		*r = false
	default:
		d.Error = errors.New("invalid bool input")
	}
}

// Reads bytes of type uint16 from byte slice
func (d *ByteDecoder) ReadUint16(r *uint16) {
	if d.checkSize(Uint16Size) != nil {
		return
	}

	*r = binary.BigEndian.Uint16(d.Bytes[d.offset:])
	d.offset += Uint16Size
}

// Reads bytes of type uint32 from byte slice
func (d *ByteDecoder) ReadUint32(r *uint32) {
	if d.checkSize(Uint32Size) != nil {
		return
	}

	*r = binary.BigEndian.Uint32(d.Bytes[d.offset:])
	d.offset += Uint32Size
}

// Reads bytes of type uint64 from byte slice
func (d *ByteDecoder) ReadUint64(r *uint64) {
	if d.checkSize(Uint64Size) != nil {
		return
	}

	*r = binary.BigEndian.Uint64(d.Bytes[d.offset:])
	d.offset += Uint64Size
}

// Reads list of bytes from byte slice
func (d *ByteDecoder) ReadByteList(r *[][]byte, size int) {
	var lsize, i uint32
	var bytes []byte
	d.ReadUint32(&lsize)

	for ; i < lsize; i += 1 {
		d.ReadBytes(&bytes, size)
		*r = append(*r, bytes)
	}
}

// Reset offset of ByteDecoder
func (d *ByteDecoder) Reset() {
	d.offset = 0
	d.Error = nil
}
