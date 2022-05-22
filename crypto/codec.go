package crypto

import (
	"encoding/hex"
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
