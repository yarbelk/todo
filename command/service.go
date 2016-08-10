package command

import "github.com/micro/go-micro"

var ServiceName string = "todo.command"

func NewService() micro.Service {
	var service micro.Service
	service = micro.NewService(
		micro.Name(ServiceName),
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
