package store

import (
	"log"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

var ServiceName string = "todo.store"

func NewService() micro.Service {
	var ss *StoreService = &StoreService{}
	var service micro.Service
	service = micro.NewService(
		micro.Name(ServiceName),
		micro.Flags(cli.StringFlag{
			Name:   "db-file",
			EnvVar: "TODO_DB_FILE",
			Value:  "todo.db",
			Usage:  "`FILENAME` is the file to use for the boltdb backed storage",
		}),
		micro.Action(func(ctx *cli.Context) {
			var err error
			ss, err = New(ctx.String("db-file"))
			if err != nil {
				panic(err.Error())
			}
			if err := service.Run(); err != nil {
				log.Fatalf(err.Error())
			}
			RegisterStorerHandler(service.Server(), ss)
			service.Run()
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
