package core

import (
	"fmt"

	"github.com/rajatxs/cosmic/codec"
	"github.com/rajatxs/cosmic/ctype"
)

const (
	TX_TRANSFER byte = iota + 1
)

type Transaction struct {
	Sequence    uint16        `json:"sequence"`
	Type        byte          `json:"type"`
	Round       uint32        `json:"round"`
	Sender      ctype.Address `json:"sender"`
	SuppliedGas uint64        `json:"suppliedGas"`
	GasLimit    uint64        `json:"gasLimit"`
	Value       uint64        `json:"value"`
	Receiver    ctype.Address `json:"receiver"`
	Proof       []byte        `json:"proof"`
	Timestamp   uint64        `json:"timestamp"`
	Expiration  uint64        `json:"expiration"`
	Payload     []byte        `json:"payload"`
}

func NewTransaction(
	txType byte,
	round uint32,
	gas uint64,
	gasLimit uint64,
	value uint64,
	receiver ctype.Address,
	proof []byte,
	expireAt uint64,
	payload []byte,
) *Transaction {
	return &Transaction{
		Type:        txType,
		Round:       round,
		Receiver:    receiver,
		Value:       value,
		SuppliedGas: gas,
		GasLimit:    gasLimit,
		Proof:       proof,
		Expiration:  expireAt,
		Payload:     payload,
	}
}

func (tx *Transaction) SanityCheck() error {
	switch {
	case tx.Sequence == 0:
		return fmt.Errorf("incorrect transaction sequence %d", tx.Sequence)
	case tx.Type != TX_TRANSFER:
		return fmt.Errorf("incorrect transaction type %d", tx.Type)
	case tx.Round == 0:
		return fmt.Errorf("transaction round should be above 0")
	case len(tx.Sender) == 0:
		return fmt.Errorf("invalid sender address %v", tx.Sender)
	case tx.GasLimit < 1:
		return fmt.Errorf("transaction gas limit should be above 0")
	case len(tx.Proof) == 0:
		return fmt.Errorf("transaction proof should not be empty")
	case tx.Timestamp < 1:
		return fmt.Errorf("invalid transaction timestamp %d", tx.Timestamp)
	}
	return nil
}

func (tx *Transaction) Marshal(r *[]byte) {
	var size int = 47 + 64 + len(tx.Payload)
	enc := codec.NewByteEncoder(size)

	enc.WriteUint16(tx.Sequence)
	enc.WriteSingleByte(tx.Type)
	enc.WriteUint32(tx.Round)
	enc.WriteBytes(tx.Sender)
	enc.WriteUint64(tx.SuppliedGas)
	enc.WriteUint64(tx.GasLimit)
	enc.WriteUint64(tx.Value)
	enc.WriteBytes(tx.Receiver)
	enc.WriteSizedBytes(tx.Proof, 64)
	enc.WriteUint64(tx.Timestamp)
	enc.WriteUint64(tx.Expiration)
	enc.WriteSizedBytes(tx.Payload, 128)

	*r = enc.Bytes
}
