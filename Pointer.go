package main

import "fmt"

func main() {
	//指针声明
	var a *int
	var b = 3
	//指针赋值
	c := &b
	a = &b
	//指针修改值
	*c = 4
	//指针获取值
	d := *c
	fmt.Printf("%p,%p,%d", a, c, d)
	//没有指针运算
}
