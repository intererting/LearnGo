package main

import "fmt"

type TestStruct struct {
	Name string
}

func TestRef() {
	a := TestStruct{
		"yuliyang",
	}
	//赋值即拷贝
	b := a
	b.Name = "changed"
	fmt.Println(a.Name)
	fmt.Println(b.Name)

	//引用类型除外
	c := map[string]string{"name": "yuliyang"}
	d := c
	d["name"] = "changed"
	fmt.Println(c["name"])
	fmt.Println(d["name"])
}
