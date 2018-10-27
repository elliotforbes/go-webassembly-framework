package router

import (
	"syscall/js"
)

type Router struct {
	Routes map[string]func() string
}

var router Router

func init() {
	router.Routes = make(map[string]func() string)
}

func NewRouter() {
	js.Global().Set("Link", js.NewCallback(Link))
	js.Global().Get("document").Call("getElementById", "view").Set("innerHTML", "")
}

func RegisterRoute(path string, component func() string) {
	router.Routes[path] = component
}

func AllRoutes() {
	println(router.Routes)
}

func Link(i []js.Value) {
	println("Link Hit")

	inner := router.Routes[i[0].String()]
	html := inner()
	js.Global().Get("document").Call("getElementById", "view").Set("innerHTML", html)
}
