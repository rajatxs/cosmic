package core

import "github.com/rajatxs/cosmic/ctype"

type AccountState struct {
	address ctype.AccountAddress
	round   uint32
	balance uint64
}

func NewAccountState() *AccountState {
	return &AccountState{
		address: []byte{},
		round:   0,
		balance: 0,
	}
}

func (as *AccountState) SetAddress(addr ctype.AccountAddress) {
	as.address = addr
}

func (as *AccountState) SetRound(val uint32) {
	as.round = val
}

func (as *AccountState) SetBalance(val uint64) {
	as.balance = val
}
