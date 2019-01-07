package router

import (
	"syscall/js"

	"github.com/elliotforbes/go-webassembly-framework/pkg/component"
)

type Router struct {
	Routes map[string]component.Component
}

var router Router

func init() {
	router.Routes = make(map[string]component.Component)
}

func NewRouter() {
	js.Global().Set("Link", js.NewCallback(Link))
	js.Global().Get("document").Call("getElementById", "view").Set("innerHTML", "")
}

func RegisterRoute(path string, component component.Component) {
	router.Routes[path] = component
}

func Link(i []js.Value) {
	println("Link Hit")

	comp := router.Routes[i[0].String()]
	html := comp.Render()

	js.Global().Get("document").Call("getElementById", "view").Set("innerHTML", html)
}
