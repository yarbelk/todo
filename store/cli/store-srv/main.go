package main

import store "github.com/yarbelk/todo/store"

// create a new command object to use in the service
func main() {
	service := store.NewService()
	service.Init()
}
