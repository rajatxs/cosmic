package core

import (
	"bytes"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/rajatxs/cosmic/codec"
	"github.com/rajatxs/cosmic/crypto"
	"github.com/rajatxs/cosmic/ctype"
	"github.com/rajatxs/cosmic/logger"
)

type BlockHeader struct {
	Id              uint64          `json:"id"`
	Height          uint64          `json:"height"`
	Version         uint16          `json:"version"`
	GasUsed         uint64          `json:"gasUsed"`
	Reward          uint64          `json:"reward"`
	TotalTx         uint16          `json:"totalTransactions"`
	StateCode       ctype.HashCode  `json:"stateCode"`
	TxCode          ctype.HashCode  `json:"txCode"`
	ParentBlockCode ctype.BlockCode `json:"parentBlockCode"`
	Time            uint64          `json:"time"`
}

func (bh *BlockHeader) IsEmpty() bool {
	return (bh.Id == 0 ||
		bh.Version == 0 ||
		bh.Height == 0 ||
		len(bh.ParentBlockCode) == 0 ||
		len(bh.StateCode) == 0 ||
		len(bh.TxCode) == 0 ||
		bh.Time < 1)
}

func NewBlockHeader(id uint64) BlockHeader {
	return BlockHeader{
		Id:              id,
		Version:         0,
		Height:          id - 1,
		ParentBlockCode: codec.NilSha256Bytes,
		StateCode:       codec.NilSha256Bytes,
		TxCode:          codec.NilSha256Bytes,
		GasUsed:         0,
		Time:            0,
	}
}

func (bh *BlockHeader) EncodeRLP() []byte {
	encoded, encodeError := rlp.EncodeToBytes(bh)

	if encodeError != nil {
		logger.Err(encodeError)
	}

	return encoded
}

func (bh *BlockHeader) Encode(r *[]byte) error {
	enc := codec.NewByteEncoder(140)

	enc.WriteUint64(bh.Id)
	enc.WriteUint64(bh.Height)
	enc.WriteUint16(bh.Version)
	enc.WriteUint64(bh.GasUsed)
	enc.WriteUint64(bh.Reward)
	enc.WriteUint16(bh.TotalTx)
	enc.WriteBytes(bh.StateCode[:])
	enc.WriteBytes(bh.TxCode[:])
	enc.WriteBytes(bh.ParentBlockCode[:])
	enc.WriteUint64(bh.Time)

	if enc.Error != nil {
		return enc.Error
	}

	*r = enc.Bytes

	return nil
}

func (bh *BlockHeader) VerifyCode(sig *[]byte) bool {
	encodedSlot := bh.EncodeRLP()
	computedSig := GenerateBlockHeaderCode(&encodedSlot)

	return bytes.Equal(*sig, computedSig)
}

func GenerateBlockHeaderCode(data *[]byte) []byte {
	return crypto.Sha256(*data)
}

func (bh *BlockHeader) DeriveCode() []byte {
	var encoded []byte = bh.EncodeRLP()
	return GenerateBlockHeaderCode(&encoded)
}

func ReadFromRLP(data []byte, bh *BlockHeader) {
	rlp.DecodeBytes(data, &bh)
}
