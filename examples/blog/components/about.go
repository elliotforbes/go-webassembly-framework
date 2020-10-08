package components

import (
	"syscall/js"

	oak "github.com/elliotforbes/go-webassembly-framework"
)

type AboutComponent struct{}

var About AboutComponent

func init() {
	oak.RegisterFunction("coolFunc", CoolFunc)
}

func CoolFunc(this js.Value, i []js.Value) interface{} {
	println("does stuff")
	return nil
}

func (a AboutComponent) Render() string {
	return `<div>
				<h2>About Component Actually Works</h2>
				<button onClick="coolFunc();">Cool Func</button>
			</div>`
}
