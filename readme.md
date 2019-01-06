Oak - The Go WebAssembly Framework
===================================

![Travis Build Status](https://api.travis-ci.org/elliotforbes/oak.svg?branch=master)


<img align="left" alt="Oak Framework" src="logo.jpg" /> With the advent of Go supporting WebAssembly, I thought I'd take a crack at building a really simple Go based WebAssembly framework that allows you to build simple frontend applications in Go, without having to dive too deep into the bushes.

## Goals

* Easier frontend application development using Go

## Tutorial

A tutorial describing Oak is avaiable here: 
https://tutorialedge.net/golang/writing-frontend-web-framework-webassembly-go/

## CLI

If you want to easily run the example in this project, I suggest you try out the new `Oak CLI` which attempts to simplify the task of writing WebAssembly applications in Go.

```s
$ make build-cli
$ cd examples/blog
$ ./oak start
Starting Server
2019/01/06 12:00:37 listening on ":8080"...
```

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

## Components 

```go
package components

import (
	"syscall/js"

	"github.com/elliotforbes/oak"
)

type AboutComponent struct{}

var About AboutComponent

func init() {
	oak.RegisterFunction("coolFunc", CoolFunc)
}

func CoolFunc(i []js.Value) {
	println("does stuff")
}

func (a AboutComponent) Render() string {
	return `<div>
						<h2>About Component Actually Works</h2>
						<button onClick="coolFunc();">Cool Func</button>
					</div>`
}
```

## Routing

```go
package main

import (
	"github.com/elliotforbes/oak"
	"github.com/elliotforbes/oak/router"

	"github.com/elliotforbes/oak/examples/blog/components"
)

func main() {
	// Starts the Oak framework
	oak.Start()

	// Starts our Router
	router.NewRouter()
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
