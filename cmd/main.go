package main

import (
	"github.com/BlenderistDev/backupmanager/internal/command"
)

func main() {
	cmd, err := command.GetCommand()
	if err != nil {
		panic(err)
	}

	err = cmd.Execute()
	if err != nil {
		panic(err)
	}
}
