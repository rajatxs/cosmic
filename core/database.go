package core

import (
	"github.com/boltdb/bolt"
	"github.com/rajatxs/cosmic/logger"
)

var Db *bolt.DB

func init() {
	conn, err := bolt.Open("cosmic.db", 0600, bolt.DefaultOptions)

	if err != nil {
		logger.Err("Couldn't connect to database", err)
		return
	}

	Db = conn
}
