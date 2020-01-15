package db

import (
	"fmt"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

func InitDB(opts *bolt.Options) error {
	err := openDB(opts)
	if err != nil {
		return err
	}
	err = createBucket()
	if err != nil {
		return err
	}
	return nil
}

func createBucket() error {
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("tasksBucket"))
		if err != nil {
			return fmt.Errorf("Bucket creation failed with error %s", err)
		}
		return nil
	})
}

func openDB(opts *bolt.Options) error {
	db, err := bolt.Open("task.db", 0600, nil)
	_ = db
	if err != nil {
		return fmt.Errorf("Failed to open db with error %s", err)
	}
	return nil
}

func CloseDB() {
	db.Close()
}
