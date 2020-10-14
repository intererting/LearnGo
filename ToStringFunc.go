package main

import "fmt"

func main() {
	a := toStringStruct{"yuliyang"}
	//调用toString方法
	fmt.Println(a)
}

type toStringStruct struct {
	name string
}

/**
toString方法的签名为String() string
*/
func (t toStringStruct) String() string {
	return t.name + " toString"
}
