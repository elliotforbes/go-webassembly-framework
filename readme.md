Oak - The Go WebAssembly Framework
===================================

With the advent of Go supporting WebAssembly, I thought I'd take a crack at building a really simple Go based WebAssembly framework that allows you to build simple frontend applications in Go, without having to dive too deep into the bushes.

## Goals

* Easier frontend application development using Go

## Todo

* Allow registration of interfaces?
* Implement a global store which can be queried/displayed easily

## Tutorial

A tutorial describing Oak is avaiable here: 
https://tutorialedge.net/golang/writing-frontend-web-framework-webassembly-go/

## Simple Example

Let's take a look at how this framework could be used in a very simple example. We'll be create a really simple app that features on function, `mycoolfunc()`. We'll kick off our Oak framework within our `main()` function and then we'll register our `coolfunc()` function.

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

We can then call our `coolfunc()` function from our `index.html` like so: 

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

## Routing

```go
package main

import (
	"github.com/elliotforbes/oak"
	"github.com/elliotforbes/oak/router"
)

func homeComponent() string {
	return "<h2>Home Component</h2>"
}

func aboutComponent() string {
	return "<h2>About Component</h2>"
}

func main() {
	// Starts the Oak framework
	oak.Start()

	// Starts our Router
	router.NewRouter()
	router.RegisterRoute("home", homeComponent)
	router.RegisterRoute("about", aboutComponent)

	// keeps our app running
	done := make(chan struct{}, 0)
	<-done
}
```

```html
<!doctype html>
<html>

<head>
	<meta charset="utf-8">
	<title>Blog</title>
	<link rel="stylesheet" href="./static/bootstrap.css">
	<link rel="stylesheet" href="./static/style.css">
	<script src="./static/wasm_exec.js"></script>
	<script src="./static/entrypoint.js"></script>
</head>
<body>	

  <div class="container">
    <h1>A Simple Blog</h1>

    <div id="view"></div>

    <button onClick="Link('home')">Home</button>
    <button onClick="Link('about')">About</button>

  </div>
</body>

</html>
```
