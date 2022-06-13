package crypto

import (
	"crypto/sha256"

	"github.com/rajatxs/cosmic/codec"
)

/** Generates SHA256 Hash of `data` */
func Sha256(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}

// Generates SHA256 hash in Hex format
func Sha256InHex(data []byte) string {
	return codec.BytesToHex(Sha256(data))
}

// Generates SHA256 hash Base58 format
func Sha256ToBase58(data []byte) string {
	return codec.EncodeBase58(Sha256(data))
}
