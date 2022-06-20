package core

import (
	"bytes"

	"github.com/rajatxs/cosmic/codec"
	"github.com/rajatxs/cosmic/ctype"
	"github.com/rajatxs/cosmic/logger"
)

const MinimalBlockSize = 36

type Block struct {
	Header       *BlockHeader                  `json:"header"`
	Transactions map[*ctype.TxCode]Transaction `json:"transactions"`
}

type BlockBuffer struct {
	Header  *bytes.Buffer
	TxCodes *bytes.Buffer
}

func NewBlockBuffer(h []byte, tx []byte) *BlockBuffer {
	return &BlockBuffer{
		Header:  bytes.NewBuffer(h),
		TxCodes: bytes.NewBuffer(tx),
	}
}

func NewBlock() *Block {
	return &Block{
		Header: &BlockHeader{},
	}
}

func (b *Block) ByteSize() (c int) {
	c = MinimalBlockSize
	c += len(b.Header.ParentBlockCode)
	c += len(b.Header.StateCode)
	c += len(b.Header.TxCode)
	return c
}

func (b *Block) TransactionCount() int {
	return len(b.Transactions)
}

func (b *Block) Marshal(buff *BlockBuffer) {
	var (
		hbuff  []byte
		txbuff *codec.ByteEncoder
	)

	txbuff = codec.NewByteEncoder(b.TransactionCount() * 32)

	if err := b.Header.Marshal(&hbuff); err != nil {
		logger.Err("Couldn't Marshal BlockHeader", err)
	}

	for tcode := range b.Transactions {
		txbuff.WriteSizedBytes(*tcode, 32)
	}

	*buff = *NewBlockBuffer(hbuff, txbuff.Bytes)
}

func (b *Block) SanityCheck() error {
	var err error

	// reject block header
	if err = b.Header.SanityCheck(); err != nil {
		return err
	}

	// apply sanity check on all transactions
	// for txCode, tx := range b.Transactions {

	// }

	return nil
}
