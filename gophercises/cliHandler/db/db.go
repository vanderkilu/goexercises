package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var tasksBucket = []byte("TASKS")
var db *bolt.DB

type Task struct {
	Id   int
	Name string
}

func InitDB(path string) error {
	var err error
	db, err = bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(tasksBucket)
		return err
	})
}

func itob(i int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

func AddTask(taskName string) (int, error) {
	var taskId int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(tasksBucket)
		id, _ := b.NextSequence()
		taskId = int(id)

		return b.Put(itob(taskId), []byte(taskName))
	})
	if err != nil {
		return -1, err
	}
	return taskId, err
}

func DeleteTask(id int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(tasksBucket)
		return b.Delete(itob(id))
	})
}

func ListTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(tasksBucket)
		cursor := b.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			tasks = append(tasks, Task{
				Id:   btoi(k),
				Name: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}
