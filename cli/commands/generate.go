package commands

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate components, routes, and services with a simple command",
	Long:  `Generate components, routes, and services with a simple command`,
	Args:  cobra.MinimumNArgs(1),
	Run:   generateExecute,
}

func generateExecute(cmd *cobra.Command, args []string) {
	switch args[0] {
	case "component":
		color.Green("Generating a New Component")
	case "service":
		color.Green("No idea how this will work but sure")
	default:
		color.Red("%s - is not a valid object\n", args[0])
	}
}
