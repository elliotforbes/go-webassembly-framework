package http

import (
	"syscall/js"
)

func RegisterCallbacks() {
	js.Global().Set("get", js.FuncOf(GetRequest))
}

func GetRequest(this js.Value, i []js.Value) interface{} {
	// req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	// req.Header.Add("js.fetch:mode", "cors")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil
	// }

	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil
	// }
	// defer resp.Body.Close()
	return nil
}
