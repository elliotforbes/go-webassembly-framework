Oak - The Go WebAssembly Framework
===================================

## Goals

* Easier frontend application development using Go

```go
package main

import (
	"syscall/js"

	"github.com/elliotforbes/oak"
)

func mycoolfunc(i []js.Value) {
	println("My Awesome Function")
}

func main() {
	oak.Start()
	oak.RegisterFunction("coolfunc", mycoolfunc)
	// keeps our app running
	done := make(chan struct{}, 0)
	<-done
}
```

This can then be called from HTML like so:

```html
<!doctype html>
<html>

<head>
	<meta charset="utf-8">
	<title>Go wasm</title>
	<script src="./static/wasm_exec.js"></script>
	<script src="./static/entrypoint.js"></script>
</head>
<body>	
    <h2>Super Simple Example</h2>
    <button class="btn btn-primary btn-block" onClick="coolfunc();" id="subtractButton">My Cool Func</button>
</body>

</html>
```