package ledger

import (
	"database/sql"
	"fmt"

	"github.com/rajatxs/cosmic/core"
	"github.com/rajatxs/cosmic/logger"
	"github.com/rajatxs/cosmic/storage"
)

const TRANSACTION_TABLE_NAME string = "transactions"

func init() {

	// creates `transactions` table if not exists
	if !storage.CheckTableExistence(TRANSACTION_TABLE_NAME) {
		storage.ExecQuery(TRANSACTION_TABLE_NAME)
		logger.Info("Created table", TRANSACTION_TABLE_NAME)
	}
}

func scanTransaction(row *sql.Row, tx *core.Transaction) error {
	var err error = row.Scan(
		&tx.Sig, &tx.Sequence, &tx.Type,
		&tx.Nonce, &tx.Sender, &tx.SuppliedGas,
		&tx.GasLimit, &tx.Value, &tx.Receiver,
		&tx.Proof, &tx.Expiration, &tx.Timestamp,
	)
	if err != nil {
		logger.Err(fmt.Sprintf("Couldn't read block header %v", tx.Sig), err)
	}

	return err
}

func scanTransactions(rows *sql.Rows, txs *[]core.Transaction) error {
	var err error

	for rows.Next() {
		tx := &core.Transaction{}

		err = rows.Scan(
			tx.Sig, tx.Sequence, tx.Type,
			tx.Nonce, tx.Sender, tx.SuppliedGas,
			tx.GasLimit, tx.Value, tx.Receiver,
			tx.Proof, tx.Expiration, tx.Timestamp,
		)

		if err != nil {
			break
		}

		*txs = append(*txs, *tx)
	}

	return err
}

// Reads Transaction by Signature
func ReadTransactionBySig(sig []byte, tx *core.Transaction) error {
	result := storage.Sql.QueryRow(
		`SELECT sig, block_id, seq, type, nonce, sender, supplied_gas, gas_limit, value, receiver, proof, expiration, ts
		FROM transactions WHERE sig = ? LIMIT 1;`,
		sig)

	return scanTransaction(result, tx)
}

// Reads All transactions by Block Id
func ReadTransactionsByBlockId(bid uint64, txs *[]core.Transaction) error {
	rows, err := storage.Sql.Query(
		`SELECT sig, block_id, seq, type, nonce, sender, supplied_gas, gas_limit, value, receiver, proof, expiration, ts
		FROM transactions WHERE block_id = ? ORDER BY seq;`,
		bid)

	if err != nil {
		logger.Err(err)
	}

	defer rows.Close()
	return scanTransactions(rows, txs)
}

// Writes new Transaction into database
func WriteTransaction(bid uint64, tx *core.Transaction) (uint64, error) {
	stmt, prepareError := storage.Sql.Prepare(`
		INSERT INTO transactions
		(sig, block_id, seq, type, nonce, sender, supplied_gas, gas_limit, value, receiver, proof, expiration, ts)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`)

	if prepareError != nil {
		logger.Err(prepareError)
	}

	defer stmt.Close()

	result, insertError := stmt.Exec(
		tx.Sig, bid, tx.Sequence,
		tx.Type, tx.Nonce, tx.Sender,
		tx.SuppliedGas, tx.GasLimit, tx.Value,
		tx.Receiver, tx.Proof, tx.Expiration,
		tx.Timestamp,
	)

	insertedId, _ := result.LastInsertId()

	return uint64(insertedId), insertError
}
