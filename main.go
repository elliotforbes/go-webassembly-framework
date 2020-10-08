package oak

import (
	"syscall/js"

	"github.com/elliotforbes/go-webassembly-framework/http"
)

func registerCallbacks() {
	http.RegisterCallbacks()
}

func RegisterFunction(functionName string, function func(this js.Value, i []js.Value) interface{}) {
	js.Global().Set(functionName, js.FuncOf(function))
}

func Start() {
	println("Oak Framework Initialized")
	registerCallbacks()
}
