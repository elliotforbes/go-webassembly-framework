package http

import (
	"syscall/js"
)

func RegisterCallbacks() {
	js.Global().Set("get", js.NewCallback(GetRequest))
}

func GetRequest(i []js.Value) {
	println("Does nothing")
}
