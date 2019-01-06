package main

import (
	"fmt"

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
	fmt.Println("Oak Command Line Interface")
	rootCmd.Execute()
}
