package core

import (
	"fmt"

	"github.com/rajatxs/cosmic/ctype"
)

const (
	TX_TRANSFER uint8 = iota
)

type Transaction struct {
	Sequence    uint16        `json:"sequence"`
	Type        uint8         `json:"type"`
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
	txType uint8,
	round uint32,
	receiver ctype.Address,
	value uint64,
	gas uint64,
	gasLimit uint64,
	expireAt uint64,
	payload []byte,
) *Transaction {
	return &Transaction{}
}

func (tx *Transaction) SanityCheck() error {
	switch {
	case tx.Sequence == 0:
		return fmt.Errorf("incorrect transaction sequence %d", tx.Sequence)
	case tx.Type != TX_TRANSFER:
		return fmt.Errorf("incorrect transaction type %d", tx.Type)
	case tx.Round == 0:
		return fmt.Errorf("transaction round should be above 0")
	case len(tx.Sender) > 0:
		return fmt.Errorf("invalid sender address %v", tx.Sender)
	case tx.GasLimit > 0:
		return fmt.Errorf("transaction gas limit should be above 0")
	case len(tx.Proof) > 0:
		return fmt.Errorf("transaction proof should not be empty")
	case tx.Timestamp > 0:
		return fmt.Errorf("invalid transaction timestamp %d", tx.Timestamp)
	}
	return nil
}
