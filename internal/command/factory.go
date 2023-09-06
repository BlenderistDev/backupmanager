package command

import (
	"fmt"
	"os"

	"github.com/BlenderistDev/backupmanager/internal/tools"
)

type Cmd interface {
	Execute() error
}

func GetCommand() (Cmd, error) {
	args, err := tools.ParseArgs(os.Args)
	if err != nil {
		return nil, err
	}
	switch args.Command {
	case "replace":
		return replace{args: args.Params}, nil
	case "delete":
		return del{args: args.Params}, nil
	case "emptydir":
		return emptydir{args: args.Params}, nil
	case "serve":
		return serve{args: args.Params}, nil
	}

	return nil, fmt.Errorf("command %s not found", args.Command)
}
