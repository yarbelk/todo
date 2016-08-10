package query

import "github.com/micro/go-micro"

var ServiceName string = "todo.query"

func NewService() micro.Service {
	var service micro.Service
	service = micro.NewService(
		micro.Name(ServiceName),
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
