package main

import "fmt"

//高阶函数
func main() {
	//函数引用
	//fmt.Println(complexFunc(math.Min))
	fmt.Println(complexFunc(minus))
}

func minus(a float64, b float64) float64 {
	return a - b
}

//高阶函数声明
func complexFunc(minus func(float64, float64) float64) float64 {
	return minus(4, 1)
}
