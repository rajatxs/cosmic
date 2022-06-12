package core

import "github.com/rajatxs/cosmic/ctype"

const (
	TX_TRANSFER uint8 = iota
)

type Transaction struct {
	Code        ctype.TxCode `json:"code"`
	Sequence    uint16       `json:"sequence"`
	Type        uint8        `json:"type"`
	Nonce       uint32       `json:"nonce"`
	Sender      []byte       `json:"sender"`
	SuppliedGas uint64       `json:"suppliedGas"`
	GasLimit    uint64       `json:"gasLimit"`
	Value       uint64       `json:"value"`
	Receiver    []byte       `json:"receiver"`
	Proof       []byte       `json:"proof"`
	Timestamp   uint64       `json:"timestamp"`
	Expiration  uint64       `json:"expiration"`
}
