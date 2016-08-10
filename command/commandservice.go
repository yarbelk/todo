package command

import (
	"golang.org/x/net/context"

	"github.com/micro/go-platform/log"
	uuid "github.com/satori/go.uuid"
	"github.com/yarbelk/todo"
	"github.com/yarbelk/todo/store"
)

type Service struct {
	Store  store.StorerClient
	logger log.Logger
}

// NewTask will assign an ID to a task definition, and store it into the storage backend
func (s *Service) NewTask(ctx context.Context, create *TaskCreate, created *TaskCreated) error {
	uuid := uuid.NewV4().String()
	created.Id = uuid
	created.Description = create.Description

	taskDef := &todo.TaskDefinition{
		Id:          created.Id,
		Description: created.Description,
		Completed:   false,
	}

	out, err := s.Store.Save(ctx, taskDef)
	if err != nil {
		return err
	}
	created.Id = out.Id
	created.Description = out.Description
	return nil
}

func New() *Service {
	return &Service{
		Store: store.NewClient(),
	}
}
