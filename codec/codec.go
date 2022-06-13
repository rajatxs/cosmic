package codec

import (
	"encoding/hex"

	"github.com/btcsuite/btcutil/base58"
)

const (
	ByteSize   int = 1
	BoolSize   int = 1
	Uint16Size int = 2
	Uint32Size int = 4
	Uint64Size int = 8
)

var NilSha256Bytes = []byte{
	00, 00, 00, 00, 00,
	00, 00, 00, 00, 00,
	00, 00, 00, 00, 00,
	00, 00, 00, 00, 00,
	00, 00, 00, 00, 00,
	00, 00, 00, 00, 00,
	00, 00,
}

// Encodes byte data into Base58 string
func EncodeBase58(d []byte) string {
	return base58.Encode(d)
}

// Decodes Base58 string value into bytes
func DecodeBase58(v string) []byte {
	return base58.Decode(v)
}

// Encodes `data` into hex string
func BytesToHex(data []byte) string {
	return "0x" + hex.EncodeToString(data)
}

// Decodes `v` hex value to `bytes`
func HexToBytes(v string) []byte {
	var decoded []byte

	// remove '0x' prefix if `v` contains
	if len(v) >= 2 && v[0] == '0' && (v[1] == 'x' || v[1] == 'X') {
		v = v[2:]
	}

	if len(v)%2 == 1 {
		v = "0" + v
	}

	decoded, _ = hex.DecodeString(v)

	return decoded
}
