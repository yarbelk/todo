package store

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

var ServiceName string = "todo.store"

var StorageFile string
var LogstashAddress string
var LogstashPort int

func NewService() micro.Service {
	var service micro.Service
	service = micro.NewService(
		micro.Name(ServiceName),
		micro.Flags(cli.StringFlag{
			Name:        "db-file",
			EnvVar:      "TODO_DB_FILE",
			Value:       "/tmp/todo.db",
			Usage:       "the file to use for the boltdb backed storage",
			Destination: &StorageFile,
		}),
		micro.Flags(cli.StringFlag{
			Name:        "logstash-address",
			EnvVar:      "TODO_LOGSTASH_ADDRESS",
			Value:       "logstash",
			Usage:       "the IP or hostname of the logstash server",
			Destination: &LogstashAddress,
		}),
		micro.Flags(cli.IntFlag{
			Name:        "logstash-port",
			EnvVar:      "TODO_LOGSTASH_PORT",
			Value:       5001,
			Usage:       "the logstash port setup to parse json",
			Destination: &LogstashPort,
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
