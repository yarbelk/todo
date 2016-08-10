package query

import "github.com/micro/go-micro"

var ServiceName string = "todo.query"

func NewService() (micro.Service, error) {
	opts := []micro.Option{
		micro.Name(ServiceName),
	}
	service := micro.NewService(opts...)

	service.Init()

	queryer := New()

	RegisterQueryerHandler(service.Server(), queryer)
	return service, nil
}

func NewClient() QueryerClient {
	opts := []micro.Option{
		micro.Name(ServiceName + ".client"),
	}
	service := micro.NewService(opts...)
	return NewQueryerClient(ServiceName, service.Client())
}
