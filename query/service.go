package query

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

var ServiceName string = "todo.query"

func NewService() micro.Service {
	var service micro.Service
	service = micro.NewService(
		micro.Name(ServiceName),
		micro.Action(func(ctx *cli.Context) {
			queryer := New()
			RegisterQueryerHandler(service.Server(), queryer)
			service.Run()
		}),
	)
	return service
}

func NewClient() QueryerClient {
	opts := []micro.Option{
		micro.Name(ServiceName + ".client"),
	}
	service := micro.NewService(opts...)
	return NewQueryerClient(ServiceName, service.Client())
}
