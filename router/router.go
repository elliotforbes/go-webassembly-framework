package router

import (
	"syscall/js"
)

type Router struct {
	Routes map[string]string
}

var router Router

func NewRouter() {
	js.Global().Set("Link", js.NewCallback(Link))
	js.Global().Get("document").Call("getElementById", "view").Set("innerHTML", "")
}

func RegisterRoute(path string, component func() string) {
	println(path)
}

func AllRoutes() {
	println(router.Routes)
}

func Link(i []js.Value) {
	println("Link Hit")

	// inner := router.Routes[i[0].String()]
	js.Global().Get("document").Call("getElementById", "view").Set("innerHTML", i[0].String())
}
