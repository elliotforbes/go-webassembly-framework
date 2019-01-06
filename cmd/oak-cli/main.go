package main

import (
	"fmt"
	"runtime"

	"github.com/elliotforbes/oak/internal/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "oak",
	Short: "Oak is a framework for building WebAssembly Applications in Go",
	Long: `A fully functioning WebAssembly Framework written in Go which allows you to create fast, 
			powerful web applications in your favorite language`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello World")
	},
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rootCmd.AddCommand(commands.StartCmd)
	rootCmd.AddCommand(commands.NewCmd)
	rootCmd.AddCommand(commands.GenerateCmd)
	rootCmd.Execute()
}
