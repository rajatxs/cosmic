package core

import "github.com/rajatxs/cosmic/ctype"

const (
	TX_TRANSFER uint8 = iota
)

type Transaction struct {
	Code        ctype.TxCode  `json:"code"`
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
