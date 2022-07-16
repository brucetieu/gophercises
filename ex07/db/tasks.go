package db

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/brucetieu/gophercises/ex07/models"
)

var DB *bolt.DB
var bucketName = "tasks"

func InitDB() {
	taskDB, err := bolt.Open("task.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Create task bucket on startup
	taskDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			log.Fatal(err)
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil

	})
	DB = taskDB
}

func CreateTask(task *models.Task) error {
	return DB.Update(func(tx *bolt.Tx) error {
		taskBucket := tx.Bucket([]byte(bucketName))
		
		// Generate next ID for task
		id, _ := taskBucket.NextSequence()
		task.ID = int(id)

		buf, err := json.Marshal(task)
		if err != nil {
			return err
		}

		return taskBucket.Put(itob(int(id)), buf)
	})
}

func DeleteTask(taskId int) (models.Task, error) {
	deletedTask, err := GetTask(taskId)
	if err != nil {
		return deletedTask, err
	}

	err = DB.Update(func(tx *bolt.Tx) error {
		taskBucket := tx.Bucket([]byte(bucketName))
		err := taskBucket.Delete(itob(taskId))
		return err
	})

	if err != nil {
		return deletedTask, err
	}

	return deletedTask, err
}

func GetTasks() ([]models.Task, error) {
	tasks := make([]models.Task, 0)

	err := DB.View(func(tx *bolt.Tx) error {
		taskBucket := tx.Bucket([]byte(bucketName))

		cur := taskBucket.Cursor()

		for k, v := cur.First(); k != nil; k, v = cur.Next() {
			var task models.Task

			_ = json.Unmarshal(v, &task)

			tasks = append(tasks, task)
		}

		return nil
	})

	return tasks, err
}

func GetTask(taskId int) (models.Task, error) {
	task := models.Task{}

	err := DB.View(func(tx *bolt.Tx) error {
		taskBucket := tx.Bucket([]byte(bucketName))
		val := taskBucket.Get(itob(taskId))

		var gotTask models.Task
		_ = json.Unmarshal(val, &gotTask)

		task = gotTask

		return nil
	})

	if task.ID == 0 {
		e := fmt.Errorf("task with ID %d does not exist", taskId)
		return models.Task{}, e
	}

	return task, err
}

func UpdateTask(taskId int) (models.Task, error) {
	updatedTask := models.Task{}
	currTask, err := GetTask(taskId)

	if err != nil {
		return updatedTask, err
	}

	err = DB.Update(func(tx *bolt.Tx) error {
		taskBucket := tx.Bucket([]byte(bucketName))
		
		if err != nil {
			return err
		}

		if currTask.Completed {
			e := fmt.Errorf("task with ID %d has already been completed", currTask.ID)
			return e
		}

		currTask.Completed = true
		updatedTask = currTask

		buf, err := json.Marshal(currTask)
		if err != nil {
			return err
		}

		err = taskBucket.Put(itob(int(currTask.ID)), buf)
		return err
	})

	if err != nil {
		return models.Task{}, err
	}

	return updatedTask, nil
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
    b := make([]byte, 8)
    binary.BigEndian.PutUint64(b, uint64(v))
    return b
}