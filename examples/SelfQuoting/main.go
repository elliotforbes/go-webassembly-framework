package main

import (
	// "fmt"
	oak "github.com/elliotforbes/go-webassembly-framework"
	"syscall/js"
)

//we set some defaults here
var disk string = "100"
var os string = "Windows"

func HandleClickDisco(this js.Value, i []js.Value) interface{} {
	sliderValue := js.Global().Get("document").Call("getElementById", "DiskSlider").Get("value")
	// fmt.Println(sliderValue)
	disk = sliderValue.String()
	QuoteResults()
	return nil
}

func HandleOsChange(this js.Value, i []js.Value) interface{} {
	osSelected := js.Global().Get("document").Call("getElementsByName", "so").Index(0).Get("checked").Bool()
	if osSelected == true {
		os = "Windows"
	} else {
		os = "Linux"
	}
	QuoteResults()
	return nil
}

//Print inside QuoteResults div the quote results.
func QuoteResults() {
	// fmt.Println("You choice a Host with " + disk + " GB of disk with " + os)
	js.Global().Get("document").Call("getElementById", "QuoteResults").Set("innerHTML", "You choice a Host with "+disk+" GB of disk with "+os)
}

func main() {
	// Starts the Oak framework
	oak.Start()

	// registers custom functions
	oak.RegisterFunction("HandleDiskElement", HandleClickDisco)
	oak.RegisterFunction("HandleOsElement", HandleOsChange)
	QuoteResults() //call the quoteresults to print the defaults

	// keeps our app running
	done := make(chan struct{}, 0)
	<-done
}
