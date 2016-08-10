package store

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

var ServiceName string = "todo.store"

var StorageFile string

func NewService() micro.Service {
	var service micro.Service
	service = micro.NewService(
		micro.Name(ServiceName),
		micro.Flags(cli.StringFlag{
			Name:        "db-file",
			EnvVar:      "TODO_DB_FILE",
			Value:       "todo.db",
			Usage:       "`FILENAME` is the file to use for the boltdb backed storage",
			Destination: &StorageFile,
		}),
	)
	return service
}

func NewClient() StorerClient {
	opts := []micro.Option{
		micro.Name(ServiceName + ".client"),
	}
	service := micro.NewService(opts...)
	return NewStorerClient(ServiceName, service.Client())
}
