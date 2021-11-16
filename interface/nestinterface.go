package main

import "fmt"

func nestedInterface() {

	//MyStruce实现了inner和outer所以具有多态
	var inter OuterInterface = MyStruce{}
	inter.inner()
	inter.outer()
}

type MyStruce struct {
}

func (myStruce MyStruce) outer() {
	fmt.Println("outer")
}

func (myStruce MyStruce) inner() {
	fmt.Println("inner")
}

type OuterInterface interface {
	//接口嵌套，可以当成组合（优于继承）
	InnerInterface
	outer()
}

type InnerInterface interface {
	inner()
}
