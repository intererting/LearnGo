package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	testStructWithJson()
}

type student struct {
	//如果要用json那么必须首字母大写，name不行
	Name string `json:"name"`
}

func testStructWithJson() {
	s := student{"yuliyang"}
	result, err := json.Marshal(s)
	if err == nil {
		jsonString := string(result)
		fmt.Println(jsonString)
	}
}

func testStringWithJson() {
	var jsonData = `{"name":"yuliyang"}`
	var a map[string]string
	err := json.Unmarshal([]byte(jsonData), &a)
	if err != nil {
		return
	}
	fmt.Println(a["name"])

	//a := make(map[string]string)
	//a["name"] = "yuliyang"
	//b, err := json.Marshal(a)
	//if nil != err {
	//	return
	//}
	//fmt.Println(string(b))
}
