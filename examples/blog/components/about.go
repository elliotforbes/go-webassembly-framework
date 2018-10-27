package components

import "github.com/elliotforbes/oak"

type AboutComponent struct{}

var About AboutComponent

func init() {
	oak.RegisterFunction("coolFunc", CoolFunc)
}

func CoolFunc() {
	println("does stuff")
}

func (a AboutComponent) Render() string {
	return `<div>
						<h2>About Component Actually Works</h2>
						<button onClick="coolFunc();">Cool Func</button>
					</div>`
}
