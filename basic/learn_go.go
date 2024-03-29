package main

import (
	"bytes"
	"fmt"
	"runtime"
)

/**
1.var用于声明变量,类型在最后,可以省略,var通常用来声明带默认初始值的属性
2.:=替代var可以用来快捷赋值,:=可以通过右值的类型自动推断左值的类型，但是只能用于函数内
3.基本类型bool string int uint float32 float64 byte rune(一个字符)
4.const声明一个常量
*/
func testBasic() {
	//basic()

	//forLoops()

	//whileLoopWithFor()

	//ifStatement()

	//switchStatement()

	//testRune()
}

func testRune() {
	a := "a中国"
	//fmt.Println(len([]byte(a)))              //7
	//fmt.Println(len(bytes.Runes([]byte(a)))) //3
	for _, r := range bytes.Runes([]byte(a)) {
		fmt.Println(string(r))
	}
}

func basic() {
	var a, b int = 3, 4
	//a, b := 3, 4
	fmt.Println(a, b)
	c := 4
	fmt.Print(c)

	result, status := add(3, 4, "success")
	fmt.Print(result, "  ", status)

	//const a = 3
	fmt.Print(convert(1.3))
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
if可以不用(),而且if表达式前可以用简单的语句，语句的作用域在if以内,语句之间用；分割
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
	//for {
	//	fmt.Println("xx")
	//}
	//i := 10
	//for i--; i > 0 {
	//	fmt.Println("%d", i)
	//}
}

//FPrintf的使用
func testFPrintf() {
	var a bytes.Buffer

	len, err := fmt.Fprintf(&a, "%d  %s", 1, "haha")
	if err != nil {
		return
	}
	fmt.Println(len)
	fmt.Println(a.String())
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
	return float64(x)
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
