package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "this starts a live-server which runs on localhost:8080",
	Long: `This runs a local development server which serves your WebAssembly application
			running on port 8080. This will automatically watch for changes to your .go source
			files and automatically recompile the lib.wasm file every time a change is made
			`,
	Run: runServer,
}

func runServer(cmd *cobra.Command, args []string) {
	fmt.Println("Starting Server")
}
