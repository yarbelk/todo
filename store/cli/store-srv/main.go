package main

import (
	"fmt"

	"github.com/micro/go-platform/log"
	tcpOutput "github.com/yarbelk/todo/log"
	"github.com/yarbelk/todo/store"
)

// create a new command object to use in the service
func main() {
	service := store.NewService()
	service.Init()
	logOutput := tcpOutput.NewOutput(
		tcpOutput.RemoteAddress("logstash"),
		tcpOutput.RemotePort(5001),
	)
	logger := log.NewLog(
		log.WithOutput(logOutput),
		log.WithLevel(log.InfoLevel),
	)
	if err := logger.Init(); err != nil {
		fmt.Printf("Panicing due to init error: ", err.Error())
		panic(err.Error())
	}

	fmt.Printf("created a logger", logger)
	logger.Infof("Created a Logger")

	ss, err := store.New(store.StorageFile, logger)
	// might faile to get the bucket, explode here if that happens
	if err != nil {
		panic(err.Error())
	}
	store.RegisterStorerHandler(service.Server(), ss)
	service.Run()
}
