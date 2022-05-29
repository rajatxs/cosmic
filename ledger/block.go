package ledger

import (
	"fmt"

	"github.com/rajatxs/cosmic/core"
	"github.com/rajatxs/cosmic/logger"
	"github.com/rajatxs/cosmic/storage"
)

const BlockHeaderTableName = "block_headers"

func init() {

	// creates `block_headers` table if not exists
	if !storage.CheckTableExistence(BlockHeaderTableName) {
		storage.ExecQuery(BlockHeaderTableName)
		logger.Info("Created table", BlockHeaderTableName)
	}
}

// Returns `BlockHeader` by `id`
func ReadBlockHeaderById(id uint64, bh *core.BlockHeader) error {
	result := storage.Sql.QueryRow(
		`SELECT id, height, version, gas_used, reward, total_tx, state_sig, tx_sig, parent_block_sig, ts
		FROM block_headers WHERE block_headers.id = ?;`,
		id)

	if scanError := result.Scan(
		&bh.Id,
		&bh.Height,
		&bh.Version,
		&bh.GasUsed,
		&bh.Reward,
		&bh.TotalTx,
		&bh.StateSig,
		&bh.TxSig,
		&bh.ParentBlockSig,
		&bh.Time,
	); scanError != nil {
		logger.Err(fmt.Sprintf("Couldn't read block header %d", id), scanError)
		return scanError
	}

	return nil
}

func GetLatestBlockId() uint64 {
	var id uint64

	row := storage.Sql.QueryRow("SELECT MAX(block_headers.id) FROM `block_headers` LIMIT 1;")
	row.Scan(&id)
	return id
}

func GetLatestBlock() *core.BlockHeader {
	var bh *core.BlockHeader
	return bh
}

// Inserts new `BlockHeader` into database
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

	encodedBlock := bh.EncodeRLP()
	blockSig = core.GenerateBlockHeaderSig(&encodedBlock)

	result, insertError := stmt.Exec(
		bh.Id,
		core.EncodeBlockHeaderSig(blockSig),
		bh.Version,
		bh.Height,
		bh.GasUsed,
		bh.Reward,
		bh.TotalTx,
		bh.EncodeStateSig(),
		bh.EncodeTxSig(),
		bh.EncodeParentBlockSig(),
		bh.Time,
	)

	insertedId, _ := result.LastInsertId()

	return uint64(insertedId), insertError
}
