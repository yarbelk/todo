package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro/cmd"
)

// create a new command object to use in the service
func newCmd() cmd.Cmd {
	command := cmd.DefaultCmd
	app := command.App()
	app.Flags = append(app.Flags,
		cli.StringFlag{
			Name:   "db-file",
			EnvVar: "TODO_DB_FILE",
			Usage:  "`FILENAME` is the path to the boltdb file used to store tasks",
			Value:  "todo.db",
		})
	return command
}
