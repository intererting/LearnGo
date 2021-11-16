package main

import "fmt"

//测试range在引用的情况下
//range如果是引用就不是副本
func testArrayRange() {
	m := map[string]string{
		"name": "yuliyang",
	}

	a := []map[string]string{
		m,
	}
	for _, value := range a {
		value["name"] = "changed"
	}
	fmt.Printf("%v", a)
}
