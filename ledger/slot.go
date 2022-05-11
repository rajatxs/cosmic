package ledger

import (
	"fmt"

	"github.com/rajatxs/cosmic/core"
	"github.com/rajatxs/cosmic/logger"
	"github.com/rajatxs/cosmic/storage"
)

func init() {
	storage.Sql.Exec(`
		CREATE TABLE IF NOT EXISTS slot_headers (
			seq INTEGER AUTOINCREMENT,
			sig VARCHAR(46) DEFAULT,
			vers SMALLINT DEFAULT 0,
			height INTEGER NOT NULL,
			gas_used INTEGER DEFAULT 0,
			state_root VARCHAR(46) NOT NULL,
			tx_root VARCHAR(46) NOT NULL,
			parent_sig VARCHAR(46) NOT NULL,
			ts INTEGER DEFAULT 0,
		
			PRIMARY KEY(seq)
		);`)
}

func GetSlotHeaderBySeq(seq uint64) {
	var seqValue uint64

	result := storage.Sql.QueryRow("SELECT * FROM slot_headers WHERE slot_headers.seq = ?;", seq)

	result.Scan(&seqValue)

	fmt.Println(seqValue)
}

func InsertSlotHeader(sh *core.SlotHeader) (uint64, error) {
	var slotSig []byte
	stmt, prepareError := storage.Sql.Prepare(`
		INSERT INTO slot_headers
		(seq, sig, vers, height, gas_used, state_root, tx_root, parent_sig, ts) VALUES
		(?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)

	if prepareError != nil {
		logger.Err(prepareError)
	}

	encodedSlot := sh.EncodeRLP()
	slotSig = core.ComputeSlotSigFromBytes(encodedSlot)

	stmt.Exec(
		sh.Sequence,
		core.EncodeSlotSig(slotSig),
		sh.Version,
		sh.Height,
		sh.GasUsed,
		sh.EncodeStateRoot(),
		sh.EncodeTxRoot(),
		sh.EncodeParentSlotSig(),
		sh.Time,
	)
	return 0, nil
}
