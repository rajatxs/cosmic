package ledger

import (
	"github.com/rajatxs/cosmic/core"
	"github.com/rajatxs/cosmic/ctype"
	"github.com/rajatxs/cosmic/logger"
	"github.com/rajatxs/cosmic/storage"
)

const AccountStateTableName string = "account_state"

func init() {

	// creates `account_state` table if not exists
	if !storage.CheckTableExistence(AccountStateTableName) {
		storage.ExecQuery(AccountStateTableName)
		logger.Info("Created table", AccountStateTableName)
	}
}

func ExistsAccountStateAddress(addr ctype.AccountAddress) bool {
	var c uint8
	result := storage.Sql.QueryRow("SELECT count(addr) FROM account_state WHERE addr = ?;", addr)
	result.Scan(&c)

	if c == 1 {
		return true
	} else {
		return false
	}
}

func UpdateAccountState(as *core.AccountState) {

}
