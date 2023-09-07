package main

import (
	"fmt"

	"github.com/BlenderistDev/backupmanager/internal/command"
)

func main() {
	err := command.GetCommand().Execute()
	if err != nil {
		fmt.Println(fmt.Sprintf("error: %v", err))
	}
}
