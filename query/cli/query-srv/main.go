package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/yarbelk/todo/query"
)

func main() {
	service := query.NewService()
	queryer := query.New()
	query.RegisterQueryerHandler(service.Server(), queryer)
	service.Init(
		micro.Name(query.ServiceName),
		micro.Action(func(ctx *cli.Context) {
			service.Run()
		}),
	)
}
