package main

import "fmt"

func main() {

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
	//接口嵌套
	InnerInterface
	outer()
}

type InnerInterface interface {
	inner()
}
