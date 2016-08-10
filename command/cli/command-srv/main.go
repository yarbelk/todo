package main

import "github.com/yarbelk/todo/command"

func main() {
	service := command.NewService()
	service.Init()
}
