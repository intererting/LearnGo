package main

import (
	"fmt"
)

func testFunc() {
	var helloInterface myHelloInterface = &myHelloStruct{
		name: "hello world",
	}
	helloInterface.sayHello()
}

type myHelloInterface interface {
	sayHello()
}

type myHelloStruct struct {
	name string
}

func (helloStruct *myHelloStruct) sayHello() {
	fmt.Println(helloStruct.name)
}

//指针实现接口和对象实现接口是两种不同的类型
//func (helloStruct myHelloStruct) sayHello() {
//	fmt.Println(helloStruct.name)
//}
