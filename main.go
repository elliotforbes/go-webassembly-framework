package oak

import (
	"syscall/js"

	"github.com/elliotforbes/go-webassembly-framework/pkg/http"
	"github.com/elliotforbes/go-webassembly-framework/pkg/utils"
)

func registerCallbacks() {
	// packages
	utils.RegisterCallbacks()
	http.RegisterCallbacks()
}

func RegisterFunction(funcName string, myfunc func(i []js.Value)) {
	js.Global().Set(funcName, js.NewCallback(myfunc))
}

func Start() {
	println("Oak Framework Initialized")
	registerCallbacks()
}
