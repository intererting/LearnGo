package main

import "fmt"

func learnMap() {
	//直接初始化
	a := map[string]string{"name": "haha"}
	//调用make初始化，如果没有调用make，那么默认为nil
	var b = make(map[string]int)
	b["age"] = 22
	fmt.Println(a, " ", b)
	//获取值，以及是否存在
	value, exits := a["name"]
	fmt.Println(value, " ", exits)
}
