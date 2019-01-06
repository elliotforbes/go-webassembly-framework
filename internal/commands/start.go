package commands

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
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
	color.Green("Starting Development Server on http://localhost:8080")
	go http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				color.Cyan("New Change Detected")

				if event.Op&fsnotify.Write == fsnotify.Write {
					color.Green("modified file: %s\n", event.Name)
				}

				out, err := exec.Command("env GOOS=js GOARCH=wasm").Output()
				if err != nil {
					fmt.Println("err: ", err)
				}
				fmt.Println(string(out[:]))
			case err := <-watcher.Errors:
				fmt.Println("Error: ", err)
			}
		}
	}()

	err = watcher.Add(".")
	if err != nil {
		log.Fatal(err)
	}

	<-done
}
