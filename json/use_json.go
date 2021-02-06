package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	//var jsonData = `{"name":"yuliyang"}`
	//var a map[string]string
	//err := json.Unmarshal([]byte(jsonData), &a)
	//if err != nil {
	//	return
	//}
	//fmt.Println(a["name"])

	a := make(map[string]string)
	a["name"] = "yuliyang"
	b, err := json.Marshal(a)
	if nil != err {
		return
	}
	fmt.Println(string(b))

}
