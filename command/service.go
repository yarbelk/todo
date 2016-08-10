package command

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

var ServiceName string = "todo.command"

func NewService() micro.Service {
	var service micro.Service
	service = micro.NewService(
		micro.Name(ServiceName),
		micro.Action(func(ctx *cli.Context) {
			commander := New()
			RegisterCommanderHandler(service.Server(), commander)
			service.Run()
		}),
	)

	return service
}

func NewClient() CommanderClient {
	opts := []micro.Option{
		micro.Name(ServiceName + ".client"),
	}
	service := micro.NewService(opts...)
	return NewCommanderClient(ServiceName, service.Client())
}
