package cli

import (
	"fmt"
	"runtime"

	"github.com/fatih/color"
	"github.com/rajatxs/cosmic/config"
)

var cmdColor *color.Color = color.New(color.FgCyan)

func ShowDefaultCommands() {
	fmt.Printf("Cosmic %s\n\n", config.CLI_VERSION)
	fmt.Printf("%s\tManage accounts\n", cmdColor.Sprint("account"))
	fmt.Printf("%s\tClear blockchain database\n", cmdColor.Sprint("clear"))
	fmt.Printf("%s\tOutputs CLI Version\n", cmdColor.Sprint("version"))
}

func ShowVersionInformation() {
	fmt.Printf("Version: %s\n", config.CLI_VERSION)
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("Arch: %s\n", runtime.GOARCH)
}

func ShowUnrecognizedCommandError(cmd *string) {
	fmt.Printf("%s: %s", color.RedString("Unrecognized command"), *cmd)
}
