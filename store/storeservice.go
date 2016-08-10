package store

import (
	"encoding/json"
	"errors"

	"github.com/yarbelk/todo"
	"golang.org/x/net/context"

	"github.com/boltdb/bolt"
)

type StoreService struct {
	Store *bolt.DB
}

func New(filename string) (ss *StoreService, err error) {
	ss = &StoreService{}
	db, err := bolt.Open(filename, 0600, nil)
	if err != nil {
		return
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("todoStore"))
		return err
	})
	ss.Store = db
	return
}

// MissingIDError is an error for when you don't have an Id, but need one
var MissingIDError = errors.New("ID field is missing")

// Save a task into the store.  This requires there to be an Id field on the input task.  So its
// an Upsert.
func (ss *StoreService) Save(ctx context.Context, in *todo.TaskDefinition, out *todo.TaskDefinition) error {
	if in.Id == "" {
		return MissingIDError
	}
	return ss.Store.Update(func(tx *bolt.Tx) error {
		*out = *in
		data, err := json.Marshal(out)
		if err != nil {
			return err
		}
		b := tx.Bucket([]byte("todoStore"))
		return b.Put([]byte(out.Id), data)
	})
}

// Load a task from the store based on its ID.
func (ss *StoreService) Load(ctx context.Context, in *todo.TaskID, out *todo.TaskDefinition) error {
	return ss.Store.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todoStore"))
		data := b.Get([]byte(in.Id))
		return json.Unmarshal(data, out)
	})
}

// All returns all of the tasks in order of their uuids.  which basically means random order
// that persists over calls.
func (ss *StoreService) All(ctx context.Context, in *todo.AllTasksParams, out *todo.TaskList) error {
	var tasks []todo.TaskDefinition
	return ss.Store.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todoStore"))
		tasks = make([]todo.TaskDefinition, 0, b.Stats().KeyN)
		return b.ForEach(func(_, taskData []byte) error {
			var task *todo.TaskDefinition = &todo.TaskDefinition{}
			if err := json.Unmarshal(taskData, task); err != nil {
				return err
			}
			tasks = append(tasks, *task)
			return nil
		})
	})
}
