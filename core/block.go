package core

import (
	"github.com/rajatxs/cosmic/ctype"
)

type Block struct {
	Header       *BlockHeader                   `json:"header"`
	Transactions *map[*ctype.TxCode]Transaction `json:"transactions"`
}

func NewBlock() *Block {
	return &Block{
		Header: &BlockHeader{},
	}
}

func (b *Block) SanityCheck() error {
	return nil
}
