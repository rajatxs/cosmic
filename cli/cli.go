package cli

import (
	"os"
)

const DEFAULT_COMMAND string = "default"

// Returns first argument as primary command
func GetCommand() string {
	if len(os.Args[1:]) > 0 {
		return os.Args[1:][0]
	}

	return DEFAULT_COMMAND
}

// Returns list of arguments of specified primary command
func GetCommandArgs() []string {
	if len(os.Args) > 2 {
		return os.Args[2:]
	}

	return []string{}
}
