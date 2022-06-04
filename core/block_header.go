package core

import (
	"bytes"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/rajatxs/cosmic/codec"
	"github.com/rajatxs/cosmic/core/ctypes"
	"github.com/rajatxs/cosmic/crypto"
	"github.com/rajatxs/cosmic/logger"
)

type BlockHeader struct {
	Id             uint64      `json:"id"`
	Sig            ctypes.Hash `json:"sig"`
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

func (bh *BlockHeader) IsEmpty() bool {
	return (bh.Id == 0 ||
		bh.Version == 0 ||
		bh.Height == 0 ||
		len(bh.ParentBlockSig) == 0 ||
		len(bh.StateSig) == 0 ||
		len(bh.TxSig) == 0 ||
		bh.Time < 1)
}

func NewBlockHeader(id uint64) BlockHeader {
	return BlockHeader{
		Id:             id,
		Version:        0,
		Height:         id - 1,
		ParentBlockSig: codec.NilSha256Bytes,
		StateSig:       codec.NilSha256Bytes,
		TxSig:          codec.NilSha256Bytes,
		GasUsed:        0,
		Time:           0,
	}
}

func (bh *BlockHeader) EncodeRLP() []byte {
	encoded, encodeError := rlp.EncodeToBytes(bh)

	if encodeError != nil {
		logger.Err(encodeError)
	}

	return encoded
}

func (bh *BlockHeader) VerifySig(sig *[]byte) bool {
	encodedSlot := bh.EncodeRLP()
	computedSig := GenerateBlockHeaderSig(&encodedSlot)

	return bytes.Equal(*sig, computedSig)
}

func GenerateBlockHeaderSig(data *[]byte) []byte {
	return crypto.Sha256(*data)
}

func (bh *BlockHeader) DeriveSig() []byte {
	var encoded []byte = bh.EncodeRLP()
	return GenerateBlockHeaderSig(&encoded)
}

func ReadFromRLP(data []byte, bh *BlockHeader) {
	rlp.DecodeBytes(data, &bh)
}
