package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {

	res, err := http.Get("http://www.baidu.com")
	if err != nil || res == nil {
		panic(err)
	}
	defer res.Body.Close()
	data, err := httputil.DumpResponse(res, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
