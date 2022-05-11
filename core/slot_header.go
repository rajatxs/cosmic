package core

import (
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/rajatxs/cosmic/core/ctypes"
	"github.com/rajatxs/cosmic/crypto"
	"github.com/rajatxs/cosmic/logger"
)

type SlotHeader struct {
	Sequence      uint64      `json:"seq"`
	Version       uint8       `json:"version"`
	Height        uint64      `json:"height"`
	ParentSlotSig ctypes.Hash `json:"parentSlotSig"`
	GasUsed       uint64      `json:"gasUsed"`
	StateRoot     ctypes.Hash `json:"stateRoot"`
	TxRoot        ctypes.Hash `json:"txRoot"`
	Time          uint64      `json:"time"`
}

func (sh *SlotHeader) IsEmpty() bool {
	return (sh.Sequence == 0 ||
		sh.Version == 0 ||
		sh.Height == 0 ||
		len(sh.ParentSlotSig) == 0 ||
		len(sh.StateRoot) == 0 ||
		len(sh.TxRoot) == 0 ||
		sh.Time < 1)
}

func NewSlotHeader(seq uint64) SlotHeader {
	return SlotHeader{
		Sequence:      seq,
		Version:       0,
		Height:        seq - 1,
		ParentSlotSig: crypto.NilSha256Bytes,
		StateRoot:     crypto.NilSha256Bytes,
		TxRoot:        crypto.NilSha256Bytes,
		GasUsed:       0,
		Time:          0,
	}
}

func (sh *SlotHeader) EncodeRLP() []byte {
	encoded, encodeError := rlp.EncodeToBytes(sh)

	if encodeError != nil {
		logger.Err(encodeError)
	}

	return encoded
}

func (sh *SlotHeader) Verify(sig []byte) bool {
	// if sh.IsEmpty() == true {
	// 	return false
	// }

	return true
}

func ComputeSlotSigFromBytes(data []byte) []byte {
	return crypto.Sha256(data)
}

func EncodeSlotSig(data []byte) string {
	return crypto.EncodeToHex(data)
}

func (sh *SlotHeader) EncodeStateRoot() string {
	return crypto.EncodeToHex(sh.StateRoot)
}

func (sh *SlotHeader) EncodeTxRoot() string {
	return crypto.EncodeToHex(sh.TxRoot)
}

func (sh *SlotHeader) EncodeParentSlotSig() string {
	return crypto.EncodeToHex(sh.ParentSlotSig)
}

func ReadFromRLP(data []byte, sh *SlotHeader) {
	rlp.DecodeBytes(data, &sh)
}
