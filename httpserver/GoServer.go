package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/test/",
		errorHandler(HandlerFileRequest))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
