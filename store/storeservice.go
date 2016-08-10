package store

import (
	"encoding/json"

	"github.com/satori/go.uuid"
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

func (ss *StoreService) Save(ctx context.Context, in *TaskDefinition, out *TaskDefinition) error {
	return ss.Store.Update(func(tx *bolt.Tx) error {
		*out = *in
		out.Id = uuid.NewV4().String()
		data, err := json.Marshal(out)
		if err != nil {
			return err
		}
		b := tx.Bucket([]byte("todoStore"))
		return b.Put([]byte(out.Id), data)
	})
}

func (ss *StoreService) Load(ctx context.Context, in *TaskID, out *TaskDefinition) error {
	return ss.Store.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todoStore"))
		data := b.Get([]byte(in.Id))
		return json.Unmarshal(data, out)
	})
}
