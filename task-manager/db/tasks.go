package db

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

var tasksBucket = []byte("tasks")
var db *bolt.DB

type Task struct {
	Key   int
	Value string
}

// Init database
func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 2 * time.Second})

	if err != nil {
		fmt.Println("Cannot open DB")
		return err
	}

	// create bucket for data
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(tasksBucket)
		return err
	})
}

// Create new task in bucket
func CreateTask(task string) (int, error) {
	var taskId int
	err := db.Update(func(tx *bolt.Tx) error {
		tasksBucket := tx.Bucket(tasksBucket)
		id, _ := tasksBucket.NextSequence()
		taskId = int(id)
		return tasksBucket.Put(itob(taskId), []byte(task))
	})
	return taskId, err
}

// List all tasks in bucket
func ListTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		tasksBucket := tx.Bucket(tasksBucket)
		tasksBucket.ForEach(func(k, v []byte) error {
			tasks = append(tasks, Task{Key: btoi(k), Value: string(v)})
			return nil
		})
		return nil
	})
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// Delete task
func DeleteTask(taskId int) error {
	return db.Update(func(tx *bolt.Tx) error {
		tasksBucket := tx.Bucket(tasksBucket)
		return tasksBucket.Delete(itob(taskId))
	})
}

// convert int to byte slice
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// convert byte slice to int
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
