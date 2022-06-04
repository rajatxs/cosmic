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

func Sha256ToHex(data []byte) string {
	return codec.BytesToHex(Sha256(data))
}
