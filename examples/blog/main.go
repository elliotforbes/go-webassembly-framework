package main

import (
	"github.com/elliotforbes/go-webassembly-framework"
	"github.com/elliotforbes/go-webassembly-framework/examples/blog/components"
	"github.com/elliotforbes/go-webassembly-framework/router"
)

func main() {
	// Starts the Oak framework
	oak.Start()

	// Starts our Router
	router.NewRouter()
	router.RegisterRoute("home", components.Home)
	router.RegisterRoute("about", components.About)

	// keeps our app running
	done := make(chan struct{}, 0)
	<-done
}
