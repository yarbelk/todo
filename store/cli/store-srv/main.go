package main

import "github.com/yarbelk/todo/store"

// create a new command object to use in the service
func main() {
	service := store.NewService()
	service.Init()
	ss, err := store.New(store.StorageFile)
	// might faile to get the bucket, explode here if that happens
	if err != nil {
		panic(err.Error())
	}
	store.RegisterStorerHandler(service.Server(), ss)
	service.Run()
}
