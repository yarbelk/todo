package command

import (
	"golang.org/x/net/context"

	uuid "github.com/satori/go.uuid"
	"github.com/yarbelk/todo/command"
	"github.com/yarbelk/todo/store"
)

type Service struct {
	Store store.StorerClient
}

func (s *Service) NewTask(ctx context.Context, create *command.TaskCreate, created *command.TaskCreated) error {
	uuid := uuid.NewV4().String()
	created.Id = uuid
	created.Description = create.Description

	taskDef := &store.TaskDefinition{
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
