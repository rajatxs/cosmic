package main

import (
	"github.com/rajatxs/cosmic/ledger"
)

func main() {
	// sh := core.NewSlotHeader(5)
	// ledger.InsertSlotHeader(&sh)

	// fmt.Println(sh)
	ledger.GetSlotHeaderBySeq(5)
}
