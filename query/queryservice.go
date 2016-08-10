package query

import (
	"github.com/yarbelk/todo"
	"github.com/yarbelk/todo/store"
	"golang.org/x/net/context"
)

type Service struct {
	Store store.StorerClient
}

func (s *Service) GetTask(ctx context.Context, in *TaskQuery, out *todo.TaskDefinition) error {
	task, err := s.Store.Load(ctx, &todo.TaskID{in.Id})
	*out = *task
	return err
}

func (s *Service) AllTasks(ctx context.Context, in *AllTasksParams, out *todo.TaskList) error {
	tasks, err := s.Store.All(ctx, &todo.AllTasksParams{})
	*out = *tasks
	return err
}

func New() *Service {
	return &Service{
		Store: store.NewClient(),
	}
}
