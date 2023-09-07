package command

import (
	"fmt"
	"os"

	"github.com/BlenderistDev/backupmanager/internal/tools"
)

// Cmd command interface
type Cmd interface {
	// Execute command execution
	Execute() error
}

// GetCommand returns command for cli request
func GetCommand() Cmd {
	args, err := tools.ParseArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		return help{}
	}
	switch args.Command {
	case "replace":
		return replace{args: args.Params}
	case "delete":
		return del{args: args.Params}
	case "emptydir":
		return emptydir{args: args.Params}
	case "serve":
		return serve{args: args.Params}
	default:
		return help{}
	}
}
