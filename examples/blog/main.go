package main

import (
	"github.com/elliotforbes/oak"
	"github.com/elliotforbes/oak/router"
)

func homeComponent() string {
	return "<h2>Home Component</h2>"
}

func aboutComponent() string {
	return "<h2>About Component</h2>"
}

func main() {
	// Starts the Oak framework
	oak.Start()

	// Starts our Router
	router.NewRouter()
	router.RegisterRoute("home", homeComponent)
	router.RegisterRoute("about", aboutComponent)

	// keeps our app running
	done := make(chan struct{}, 0)
	<-done
}
