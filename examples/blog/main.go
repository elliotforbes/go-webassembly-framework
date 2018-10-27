package main

import (
	"syscall/js"

	"github.com/elliotforbes/oak"
	"github.com/elliotforbes/oak/examples/blog/components"
	"github.com/elliotforbes/oak/router"
)

func coolFunc(i []js.Value) {
	println("Does something cool")
}

func main() {
	// Starts the Oak framework
	oak.Start()

	oak.RegisterFunction("coolFunc", coolFunc)

	// Starts our Router
	router.NewRouter()
	router.RegisterRoute("home", components.Home)
	router.RegisterRoute("about", components.About)

	// keeps our app running
	done := make(chan struct{}, 0)
	<-done
}
