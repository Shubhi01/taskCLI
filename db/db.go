package db

import (
	"encoding/binary"
	"fmt"

	"github.com/boltdb/bolt"
)

var db *bolt.DB
var tasksBucket = []byte("tasks")

func InitDB(dbPath string, opts *bolt.Options) error {
	err := openDB(dbPath, opts)
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
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("tasksBucket"))
		if err != nil {
			return fmt.Errorf("Bucket creation failed with error %s", err)
		}
		return nil
	})

	return err
}

func openDB(dbPath string, opts *bolt.Options) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return fmt.Errorf("Failed to open db with error %s", err)
	}
	return nil
}

func CloseDB() {
	db.Close()
}

func AddTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(tasksBucket))

		id64, _ := b.NextSequence()
		id = int(id64)

		return b.Put(itob(id), []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, err
}

func ListTask() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(tasksBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			id := binary.BigEndian.Uint64(k)
			tasks = append(tasks, Task{Id: int(id), Description: string(v)})
		}

		return nil
	})

	return tasks, err
}

func DeleteTask(id int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(tasksBucket)
		err := b.Delete(itob(id))
		return err
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

type Task struct {
	Id          int
	Description string
}
