package main

import (
	"github.com/rajatxs/cosmic/cli"
)

func main() {
	cli.DispatchAction(
		cli.GetCommand(),
		cli.GetCommandArgs(),
	)
}
