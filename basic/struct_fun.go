package main

import (
	"fmt"
)

//扩展方法
func structFun() {
	//person := Person{"yuliyang"}
	//fmt.Println(person.test(3))

	//var a MyInt = -3
	//fmt.Println(a.test2())

	myPoint := MyPoint{
		x: 0,
		y: 0,
	}
	myPoint.test3()
	fmt.Println(myPoint)

}

type Person struct {
	name string
}

//最常用的结构体扩展
func (p Person) test(a int) int {
	fmt.Println(p.name)
	return a
}

type MyInt int

//对于基本类型的扩展
func (a MyInt) test2() int {
	if a < 0 {
		return int(-a)
	} else {
		return int(a)
	}
}

type MyPoint struct {
	x float64
	y float64
}

//可以给指针扩展，但是类型不能是基本类型
//指针扩展可以修改对象的值
func (myPoint *MyPoint) test3() {
	myPoint.x = 3
	myPoint.y = 4
}
