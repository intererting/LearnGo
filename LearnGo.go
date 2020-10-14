package main

import (
	"fmt"
	"runtime"
)

/**
1.var用于声明变量,类型在最后,可以省略
2.:=替代var可以用来快捷赋值,:=可以通过右值的类型自动推断左值的类型，但是只能用于函数内
3.基本类型bool string int uint float32 float64 byte
4.const声明一个常量
*/
func main() {
	//var a, b int = 3, 4
	//c := 4
	//fmt.Print(c)

	//result, status := add(3, 4, "success")
	//fmt.Print(result, "  ", status)

	//const a = 3
	//fmt.Print(convert(1.3))

	//forLoops()

	//whileLoopWithFor()

	//ifStatement()

	//switchStatement()

	fmt.Println(deferStatement())
}

/**
a 2
defer 1
before
3
*/
func deferStatement() (i int) {
	a := 1
	//defer是一个栈，FILO
	defer fmt.Println("before")
	//defer会先计算，但是会后执行，所以这里的a是1不是2
	defer fmt.Println("defer", a);
	//defer可以修改返回值，所以这里的结果为3
	defer func() { i++ }()
	a++
	fmt.Println("a", a)
	return a
}

/**
switch的基本使用
*/
func switchStatement() {
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("linux")
		//无条件执行下一个case分支
		fallthrough
	case "windows":
		fmt.Println("windows")
	case "darwin":
		fmt.Println("macos")
	default:
		fmt.Print("other")
	}
}

/**
if可以不用(),而且if表达式前可以见简单的语句，语句的作用域在if以内,语句之间用；分割
*/
func ifStatement() {
	a := 1
	if b := a + 1; b < 3 {
		fmt.Println(a, b)
	}
}

/**
用for循环实现while循环
*/
func whileLoopWithFor() {
	//i := 0
	//for i < 10 {
	//	fmt.Print(i)
	//	i++
	//}

	//等同于while(true)
	for {
		fmt.Println("xx")
	}
}

/**
go 只有for这一种循环，而且不用带()
*/
func forLoops() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

/**
没有类型的自动转化
*/
func convert(x float32) float64 {
	return float64((x))
}

/**
1.func 声明方法
2.相同类型参数可以在最后声明类型
3.返回值可以有多个
4.获取多个返回值用:=
5.返回值可以被命名,比如下面的x,y一般情况省略
*/
func add(x, y int, str string) (int, string) {
	return x + y, str
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
