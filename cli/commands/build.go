package commands

import (
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var BuildCmd = &cobra.Command{
	Use:   "build",
	Short: "builds a lightweight compressed WebAssembly binary",
	Run:   buildExecute,
}

func buildExecute(cmd *cobra.Command, args []string) {
	out, err := exec.Command("env", "tinygo", "build", "-o", "lib.wasm", "-target", "wasm", "./main.go").Output()
	if err != nil {
		color.Red("err: ", err)
	}
	color.Green(string(out[:]))
}
