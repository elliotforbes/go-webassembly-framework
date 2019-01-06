package main

import (
	"syscall/js"

	"github.com/elliotforbes/oak"
)

func mycoolfunc(i []js.Value) {
	println("My Awesome Function")
}

func testfunc(i []js.Value) {
	println("holy shit this works")
}

func main() {
	// Starts the Oak framework
	oak.Start()

	// registers custom functions
	oak.RegisterFunction("coolfunc", mycoolfunc)
	oak.RegisterFunction("holyshit", testfunc)

	// keeps our app running
	done := make(chan struct{}, 0)
	<-done
}
