package tools

import "fmt"

// Args struct for parsed cli args
type Args struct {
	Command string
	Params  map[string]string
}

// ParseArgs parse cli args to Args
// args[1] is command
// other args are map[string]string
func ParseArgs(args []string) (*Args, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("command not found")
	}

	command := args[1]

	params := make(map[string]string)

	i := 2
	for {
		var name string
		var value string
		if i < len(args) {
			name = args[i]
		} else {
			break
		}
		i++
		if i < len(args) {
			value = args[i]
		}
		i++

		if name != "" {
			params[name] = value
		}
	}

	return &Args{
		Command: command,
		Params:  params,
	}, nil
}
