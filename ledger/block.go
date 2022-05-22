package ledger

import (
	"fmt"

	"github.com/rajatxs/cosmic/core"
	"github.com/rajatxs/cosmic/crypto"
	"github.com/rajatxs/cosmic/logger"
	"github.com/rajatxs/cosmic/storage"
)

const BlockHeaderTableName = "block_headers"

func init() {

	// creates `block_headers` table if not exists
	if !storage.CheckTableExistence(BlockHeaderTableName) {
		storage.Sql.Exec(fmt.Sprintf(
			`CREATE TABLE %s ( 
				id UNSIGNED INTEGER PRIMARY KEY, 
				sig BLOB(32) NOT NULL, 
				height UNSIGNED INTEGER NOT NULL, 
				version UNSIGNED SMALLINT DEFAULT 0, 
				gas_used UNSIGNED INTEGER DEFAULT 0,
				reward UNSIGNED INTEGER DEFAULT 0,
				total_tx UNSIGNED SMALLINT DEFAULT 0, 
				state_sig BLOB(32) NOT NULL, 
				tx_sig BLOB(32) NOT NULL, 
				parent_block_sig BLOB(32) NOT NULL, 
				ts UNSIGNED INTEGER DEFAULT 0
			);`,
			BlockHeaderTableName))

		logger.Info("Created table", "block_headers")
	}
}

// Returns `BlockHeader` by `seq`
func GetBlockHeaderById(seq uint64) *core.BlockHeader {
	var stateSig, txSig, parentBlockSig string
	var sh *core.BlockHeader = &core.BlockHeader{}

	result := storage.Sql.QueryRow(
		`SELECT id, height, version, gas_used, reward, total_tx, state_tx, tx_sig, parent_block_sig, ts
		FROM block_headers WHERE block_headers.seq = ?;`,
		seq)

	result.Scan(
		&sh.Id,
		&sh.Height,
		&sh.Version,
		&sh.GasUsed,
		&sh.Reward,
		&sh.TotalTx,
		&stateSig,
		&txSig,
		&parentBlockSig,
		&sh.Time,
	)

	sh.StateSig = crypto.HexToBytes(stateSig)
	sh.TxSig = crypto.HexToBytes(txSig)
	sh.ParentBlockSig = crypto.HexToBytes(parentBlockSig)

	return sh
}

func GetLastInsertedBlockId() uint64 {
	var seq uint64

	row := storage.Sql.QueryRow("SELECT MAX(block_headers.id) FROM `block_headers` LIMIT 1;")
	row.Scan(&seq)
	return seq
}

// Inserts new `BlockHeader` into database
func WriteBlockHeader(sh *core.BlockHeader) (uint64, error) {
	var slotSig []byte

	stmt, prepareError := storage.Sql.Prepare(`
		INSERT INTO block_headers
		(id, sig, version, height, gas_used, reward, total_tx, state_sig, tx_sig, parent_block_sig, ts) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)

	if prepareError != nil {
		logger.Err(prepareError)
	}

	encodedBlock := sh.EncodeRLP()
	slotSig = core.GenerateBlockHeaderSig(&encodedBlock)

	result, insertError := stmt.Exec(
		sh.Id,
		core.EncodeBlockHeaderSig(slotSig),
		sh.Version,
		sh.Height,
		sh.GasUsed,
		sh.Reward,
		sh.TotalTx,
		sh.EncodeStateSig(),
		sh.EncodeTxSig(),
		sh.EncodeParentBlockSig(),
		sh.Time,
	)

	insertedId, _ := result.LastInsertId()

	return uint64(insertedId), insertError
}
