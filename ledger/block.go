package ledger

import (
	"database/sql"
	"fmt"

	"github.com/rajatxs/cosmic/core"
	"github.com/rajatxs/cosmic/logger"
	"github.com/rajatxs/cosmic/storage"
)

const BlockHeaderTableName string = "block_headers"

func init() {

	// creates `block_headers` table if not exists
	if !storage.CheckTableExistence(BlockHeaderTableName) {
		storage.ExecQuery(BlockHeaderTableName)
		logger.Info("Created table", BlockHeaderTableName)
	}
}

func scanBlockHeader(row *sql.Row, bh *core.BlockHeader) error {
	var err error = row.Scan(
		&bh.Id,
		&bh.Sig,
		&bh.Height,
		&bh.Version,
		&bh.GasUsed,
		&bh.Reward,
		&bh.TotalTx,
		&bh.StateSig,
		&bh.TxSig,
		&bh.ParentBlockSig,
		&bh.Time,
	)
	if err != nil {
		logger.Err(fmt.Sprintf("Couldn't read block header %d", bh.Id), err)
	}

	return err
}

// Reads `BlockHeader` by `id`
func ReadBlockHeaderById(id uint64, bh *core.BlockHeader) error {
	result := storage.Sql.QueryRow(
		`SELECT id, sig, height, version, gas_used, reward, total_tx, state_sig, tx_sig, parent_block_sig, ts
		FROM block_headers WHERE block_headers.id = ?;`,
		id)

	return scanBlockHeader(result, bh)
}

// Reads BlockHeader by Signarure
func ReadBlockHeaderBySig(sig []byte, bh *core.BlockHeader) error {
	result := storage.Sql.QueryRow(
		`SELECT id, sig, height, version, gas_used, reward, total_tx, state_sig, tx_sig, parent_block_sig, ts
		FROM block_headers WHERE block_headers.sig = ?;`,
		sig)

	return scanBlockHeader(result, bh)
}

// Reads last inserted block id
func GetLatestBlockId() uint64 {
	var id uint64

	row := storage.Sql.QueryRow("SELECT MAX(block_headers.id) FROM `block_headers` LIMIT 1;")
	row.Scan(&id)
	return id
}

// Reads last inserted BlockHeader
func ReadLatestBlockHeader(bh *core.BlockHeader) error {
	result := storage.Sql.QueryRow(
		`SELECT id, sig, height, version, gas_used, reward, total_tx, state_sig, tx_sig, parent_block_sig, ts
		FROM block_headers WHERE block_headers.id = (SELECT MAX(id) FROM block_headers) LIMIT 1;`,
	)
	return scanBlockHeader(result, bh)
}

// Writes new `BlockHeader` into database
func WriteBlockHeader(bh *core.BlockHeader) (uint64, error) {
	var blockSig []byte

	stmt, prepareError := storage.Sql.Prepare(`
		INSERT INTO block_headers
		(id, sig, version, height, gas_used, reward, total_tx, state_sig, tx_sig, parent_block_sig, ts) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`)

	if prepareError != nil {
		logger.Err(prepareError)
	}

	defer stmt.Close()

	encodedBlock := bh.EncodeRLP()
	blockSig = core.GenerateBlockHeaderSig(&encodedBlock)

	result, insertError := stmt.Exec(
		bh.Id,
		blockSig,
		bh.Version,
		bh.Height,
		bh.GasUsed,
		bh.Reward,
		bh.TotalTx,
		bh.StateSig,
		bh.TxSig,
		bh.ParentBlockSig,
		bh.Time,
	)

	insertedId, _ := result.LastInsertId()

	return uint64(insertedId), insertError
}

func WriteBlockTransactions(id uint64, txs *[]core.Transaction) {

}
