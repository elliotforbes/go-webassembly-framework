package commands

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var files = map[string]string{
	"static/entrypoint.js": "https://raw.githubusercontent.com/elliotforbes/oak/master/examples/calculator/static/entrypoint.js",
	"static/wasm_exec.js":  "https://raw.githubusercontent.com/elliotforbes/oak/master/examples/calculator/static/wasm_exec.js",
	"main.go":              "https://raw.githubusercontent.com/elliotforbes/oak/master/examples/calculator/main.go",
	"index.html":           "https://raw.githubusercontent.com/elliotforbes/oak/master/examples/calculator/index.html",
}

var NewCmd = &cobra.Command{
	Use:   "new",
	Short: "Automatically creates a new WebAssembly Project",
	Long: `Automatically generates the files and folder structure needed
			for a new WebAssembly application.`,
	Args: cobra.MinimumNArgs(1),
	Run:  newExecute,
}

func newExecute(cmd *cobra.Command, args []string) {
	color.Green("Generating New Oak WebAssembly Project: %s\n", args[0])

	err := createDirectories(args[0])
	if err != nil {
		color.Red("Error Creating new WebAssembly Project")
		os.Exit(1)
	}

	err = getFiles(args[0])
	if err != nil {
		color.Red("Error Creating new WebAssembly Project")
		os.Exit(1)
	}

}

func createDirectories(projectname string) error {
	basepath := filepath.Join(projectname)

	dirs := []string{
		filepath.Join(basepath, "components"),
		filepath.Join(basepath, "static"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0777); err != nil {
			color.Red("Error when generating new WebAssembly Project")
			color.Red("Error: %+v\n", err)
			return err
		}
		color.Green("Successfully Created Directory: %s\n", dir)
	}
	return nil
}

func getFiles(projectname string) error {
	for key, location := range files {
		resp, err := http.Get(location)
		if err != nil {
			color.Red("Error: %s\n", err)
			return err
		}
		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			color.Red("Error: %s\n", err)
			return err
		}

		err = ioutil.WriteFile(filepath.Join(projectname, key), responseData, 0777)
		if err != nil {
			color.Red("Error: %s\n", err)
			return err
		}
		color.Green("Successfully Created File: %s\n", key)

	}
	return nil
}
