package main

import (
	"syscall/js"

	"github.com/elliotforbes/go-webassembly-framework"
)

func mycoolfunc(this js.Value, i []js.Value) interface{} {
	println("My Awesome Function")
	return nil
}

func main() {
	// Starts the Oak framework
	oak.Start()

	// registers custom functions
	oak.RegisterFunction("coolfunc", mycoolfunc)

	// keeps our app running
	done := make(chan struct{}, 0)
	<-done
}
