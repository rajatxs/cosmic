package core

type Block struct {
	Header *BlockHeader `json:"header"`
	Body   *BlockBody   `json:"body"`
}

func NewBlock() *Block {
	return &Block{
		Header: &BlockHeader{},
		Body:   &BlockBody{},
	}
}
