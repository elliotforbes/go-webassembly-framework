package main

import (
	"syscall/js"

	"github.com/elliotforbes/oak"
)

func mycoolfunc(i []js.Value) {
	println("My Awesome Function")
}

func main() {
	oak.Start()
	oak.RegisterFunction("coolfunc", mycoolfunc)
	// keeps our app running
	done := make(chan struct{}, 0)
	<-done
}
