package core

import (
	"errors"

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
	switch {
	case b.Header.Id == 0:
		return errors.New("invalid Block Id")
	case b.Header.Height >= b.Header.Id:
		return errors.New("invalid Block Height")
	case b.Header.Time > 0:
		return errors.New("invalid Block Time")
	case b.Header.Version > 0:
		return errors.New("invalid Block Version")
	case len(b.Header.ParentBlockCode) != 32:
		return errors.New("invalid parent block code")
	case len(b.Header.StateCode) != 32:
		return errors.New("invalid state code")
	case len(b.Header.TxCode) != 32:
		return errors.New("invalid tx code")
	}

	return nil
}
