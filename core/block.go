package core

import (
	"github.com/rajatxs/cosmic/codec"
	"github.com/rajatxs/cosmic/crypto"
	"github.com/rajatxs/cosmic/ctype"
	"github.com/rajatxs/cosmic/logger"
)

const (
	MinimalBlockSize = 36
	MaximumBlockSize = 1036
)

type Block struct {
	Header       *BlockHeader                  `json:"header"`
	Transactions map[*ctype.TxCode]Transaction `json:"transactions"`
}

type EncodedBlock struct {
	Header  []byte
	TxCodes []byte
}

func NewEncodedBlock(h []byte, tx []byte) *EncodedBlock {
	return &EncodedBlock{
		Header:  h,
		TxCodes: tx,
	}
}

func NewBlock(h *BlockHeader) *Block {
	return &Block{
		Header:       h,
		Transactions: make(map[*ctype.TxCode]Transaction),
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

func (b *Block) margeTxCodes() ([]byte, error) {
	var enc *codec.ByteEncoder
	var size int = len(b.Transactions) * 32

	enc = codec.NewByteEncoder(size)

	for tcode := range b.Transactions {
		enc.WriteFixedBytes(*tcode, 32)

		if enc.Error != nil {
			return nil, enc.Error
		}
	}

	return enc.Bytes, nil
}

func (b *Block) EncodeBlock() (*EncodedBlock, error) {
	var (
		hbuff  []byte
		txbuff []byte
		err    error
	)

	if hbuff, err = b.Header.Marshal(); err != nil {
		logger.Err("couldn't encode BlockHeader", err)
		return nil, err
	}

	if txbuff, err = b.margeTxCodes(); err != nil {
		logger.Err("couldn't encode txcodes", err)
		return nil, err
	}

	return NewEncodedBlock(hbuff, txbuff), err
}

func (b *Block) AddTransaction(tx *Transaction) {
	var txcode []byte
	txdata, err := tx.Marshal()

	if err != nil {
		panic(err)
	}

	txcode = crypto.Sha256(txdata)
	b.Transactions[&txcode] = *tx
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
