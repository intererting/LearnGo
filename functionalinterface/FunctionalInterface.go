package main

import "fmt"

//函数式编程
func main() {
	//函数式编程标准写法
	//addr := getAddr(3)
	//var result int
	//for i := 0; i < 5; i++ {
	//	result, addr = addr(i)
	//}
	//fmt.Println(result)

	//闭包写法
	addr := getAddrClosure(3)
	for i := 0; i < 5; i++ {
		fmt.Println(addr(i))
	}
}

func getAddrClosure(base int) func(int) int {
	sum := base
	return func(v int) int {
		//通过sum这个自由变量保存中间的状态
		sum += v
		return sum
	}
}

//接受值类型  返回值类型  用于下次迭代，保存自由变量
type isAddr func(int) (int, isAddr)

func getAddr(base int) isAddr {
	return func(v int) (int, isAddr) {
		return base + v, getAddr(base + v)
	}
}
