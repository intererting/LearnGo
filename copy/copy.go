package main

import "fmt"

func main() {
	//切片拷贝，dst的结构类型必须与src相同，不然失败
	//a := []int{1}
	//b := make([]int, 1)
	//copy(b, a)
	//fmt.Printf("%v", b)

	//copy是浅拷贝
	a := map[string]string{
		"name": "yuliyang",
	}
	arr := []map[string]string{a}

	arrCopy := make([]map[string]string, 1)
	copy(arrCopy, arr)
	a["name"] = "changed"
	fmt.Printf("%v", arrCopy)

}
