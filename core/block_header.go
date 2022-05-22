package core

import (
	"bytes"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/rajatxs/cosmic/core/ctypes"
	"github.com/rajatxs/cosmic/crypto"
	"github.com/rajatxs/cosmic/logger"
)

type BlockHeader struct {
	Id             uint64      `json:"id"`
	Height         uint64      `json:"height"`
	Version        uint8       `json:"version"`
	GasUsed        uint64      `json:"gasUsed"`
	Reward         uint64      `json:"reward"`
	TotalTx        uint16      `json:"totalTransactions"`
	StateSig       ctypes.Hash `json:"stateSignature"`
	TxSig          ctypes.Hash `json:"txSignature"`
	ParentBlockSig ctypes.Hash `json:"parentBlockSig"`
	Time           uint64      `json:"time"`
}

func (sh *BlockHeader) IsEmpty() bool {
	return (sh.Id == 0 ||
		sh.Version == 0 ||
		sh.Height == 0 ||
		len(sh.ParentBlockSig) == 0 ||
		len(sh.StateSig) == 0 ||
		len(sh.TxSig) == 0 ||
		sh.Time < 1)
}

func NewBlockHeader(id uint64) BlockHeader {
	return BlockHeader{
		Id:             id,
		Version:        0,
		Height:         id - 1,
		ParentBlockSig: crypto.NilSha256Bytes,
		StateSig:       crypto.NilSha256Bytes,
		TxSig:          crypto.NilSha256Bytes,
		GasUsed:        0,
		Time:           0,
	}
}

func (sh *BlockHeader) EncodeRLP() []byte {
	encoded, encodeError := rlp.EncodeToBytes(sh)

	if encodeError != nil {
		logger.Err(encodeError)
	}

	return encoded
}

func (sh *BlockHeader) VerifySig(sig *[]byte) bool {
	encodedSlot := sh.EncodeRLP()
	computedSig := GenerateBlockHeaderSig(&encodedSlot)

	return bytes.Equal(*sig, computedSig)
}

func GenerateBlockHeaderSig(data *[]byte) []byte {
	return crypto.Sha256(*data)
}

func EncodeBlockHeaderSig(data []byte) string {
	return crypto.BytesToHex(data)
}

func (sh *BlockHeader) EncodeStateSig() string {
	return crypto.BytesToHex(sh.StateSig)
}

func (sh *BlockHeader) EncodeTxSig() string {
	return crypto.BytesToHex(sh.TxSig)
}

func (sh *BlockHeader) EncodeParentBlockSig() string {
	return crypto.BytesToHex(sh.ParentBlockSig)
}

func ReadFromRLP(data []byte, sh *BlockHeader) {
	rlp.DecodeBytes(data, &sh)
}
