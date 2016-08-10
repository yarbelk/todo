package command

import "github.com/micro/go-micro"

var ServiceName string = "todo.command"

func NewService() (micro.Service, error) {
	opts := []micro.Option{
		micro.Name(ServiceName),
	}
	service := micro.NewService(opts...)

	service.Init()

	commander := New()

	RegisterCommanderHandler(service.Server(), commander)
	return service, nil
}

func NewClient() QueryerClient {
	opts := []micro.Option{
		micro.Name(ServiceName + ".client"),
	}
	service := micro.NewService(opts...)
	return NewCommanderClient(ServiceName, service.Client())
}
