package storage

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/rajatxs/cosmic/core"
)

func init() {
	UseBucket(blockBucketId)
}

func WriteBlock(b *core.Block) {
	// var bc []byte = b.

	core.Db.Update(func(tx *bolt.Tx) error {
		buck := getBlockBucket(tx)
		return buck.Put([]byte("first"), []byte("Test Block"))
	})
}

func ReadBlock() {
	core.Db.View(func(tx *bolt.Tx) error {
		buck := getBlockBucket(tx)
		value := buck.Get([]byte("first"))

		fmt.Println("Value", string(value))
		return nil
	})
}
