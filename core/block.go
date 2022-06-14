package core

import "errors"

type Block struct {
	Header       *BlockHeader   `json:"header"`
	Transactions *[]Transaction `json:"transactions"`
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
	}

	return nil
}
