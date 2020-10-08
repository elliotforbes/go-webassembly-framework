package router

import (
	"syscall/js"

	"github.com/elliotforbes/go-webassembly-framework/component"
)

type Router struct {
	Routes map[string]component.Component
}

var router Router

func init() {
	router.Routes = make(map[string]component.Component)
}

func NewRouter() {
	js.Global().Set("Link", js.FuncOf(Link))
	js.Global().Get("document").Call("getElementById", "view").Set("innerHTML", "")
}

func RegisterRoute(path string, component component.Component) {
	router.Routes[path] = component
}

func Link(this js.Value, i []js.Value) interface{} {
	println("Link Hit")

	comp := router.Routes[i[0].String()]
	html := comp.Render()

	js.Global().Get("document").Call("getElementById", "view").Set("innerHTML", html)
	return nil
}
