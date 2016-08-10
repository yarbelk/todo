package store

import "github.com/micro/go-micro"

var ServiceName string = "todoStore"

func NewService(fileName string) (micro.Service, error) {
	opts := []micro.Option{
		micro.Name(ServiceName),
	}
	service := micro.NewService(opts...)

	service.Init()

	ss, err := New(fileName)
	if err != nil {
		return nil, err
	}

	RegisterStorerHandler(service.Server(), ss)
	return service, nil
}

func NewClient() StorerClient {
	opts := []micro.Option{
		micro.Name(ServiceName + ".client"),
	}
	service := micro.NewService(opts...)
	return NewStorerClient(ServiceName, service.Client())
}
