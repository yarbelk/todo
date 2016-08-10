package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/yarbelk/todo/command"
)

func main() {
	service := command.NewService()
	commander := command.New()
	command.RegisterCommanderHandler(service.Server(), commander)
	service.Init(
		micro.Name(command.ServiceName),
		micro.Action(func(ctx *cli.Context) {
			service.Run()
		}),
	)
}
