package cli

// Runs different function based on specified primary command
func DispatchAction(cmd string, args []string) {
	switch cmd {
	case "default":
		{
			ShowDefaultCommands()
			break
		}
	case "version":
		{
			ShowVersionInformation()
			break
		}
	default:
		{
			ShowUnrecognizedCommandError(&cmd)
			break
		}
	}
}
