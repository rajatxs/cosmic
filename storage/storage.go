package storage

import (
	"github.com/boltdb/bolt"
	"github.com/rajatxs/cosmic/core"
	"github.com/rajatxs/cosmic/logger"
)

var (
	blockBucketId = [4]byte{0xc9, 0x41, 0x90, 0x4e}
)

func getBlockBucket(tx *bolt.Tx) *bolt.Bucket {
	return tx.Bucket(blockBucketId[:])
}

func UseBucket(name [4]byte) {
	core.Db.Update(func(tx *bolt.Tx) error {
		_, e := tx.CreateBucketIfNotExists(name[:])

		if e != nil {
			logger.Err("Couldn't use bucket", name)
			return e
		}

		return nil
	})
}
