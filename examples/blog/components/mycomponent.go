package components

import (
	"github.com/elliotforbes/oak"
)

func Render() string {
	return ```
		<h2>My Awesome Component</h2>
	```
}

func SomethingCool() {
	println("This totally does something cool")
}

func RegisterFunctions() {
	oak.RegisterFunction("SomethingCool", SomethingCool)
}
