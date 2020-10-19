package main

import "fmt"

func main() {
	var a MyInterface = myInterfaceImpl{name: "yuliyang"}
	fmt.Println(a.test())
	test1(a)

	testAny(3)
	testAny(a)
}

//定义接口,这个等价于java8的FunctionalInterface
type MyInterface interface {
	test() string
}

type myInterfaceImpl struct {
	name string
}

//接口的隐式实现，方法名签名和接口相同，那么myInterfaceImpl就隐式实现了MyInterface
//duck type 接口的实际类型不用声明，只需要具备某些行为就认为是某个类型
func (m myInterfaceImpl) test() string {
	return m.name
}

//接口作为参数和返回值
func test1(m MyInterface) MyInterface {
	fmt.Printf("%T", m)
	m.test()
	return m
}

/**
interface{}可以代表任何值，在java就是Object，在kotlin就是Any
*/
func testAny(a interface{}) {
	fmt.Println("interface{}")
}
