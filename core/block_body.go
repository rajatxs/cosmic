package core

type BlockBody struct {
	Sequence     uint64  `json:"seq"`
	Transactions []uint8 `json:"transactions"`
}
