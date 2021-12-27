package main

import (
	"fmt"
	"log"
)

func main() {
	//result := deferStatement()
	//fmt.Println(result)
	result := deferChanData()
	for key, value := range result {
		fmt.Println(key, value)
	}
}

/**
a 2
2021/02/05 09:44:36 before a 2
2021/02/05 09:44:36 after a 3
defer 1
before
2021/02/05 09:44:36 result   2
*/
func deferStatement() (i int) {
	a := 1
	//defer是一个栈，FILO
	defer fmt.Println("before")
	//普通表达式defer会先计算，但是会后执行，所以这里的a是1不是2
	defer fmt.Println("defer", a)
	//defer不能修改返回值
	defer func() {
		//闭包会持有引用,但是不能修改返回值
		log.Printf("before a %d", a)
		a++
		log.Printf("after a %d", a)
	}()
	a++
	return a
}

//go语言的defer和java的finally一样都是值传递,可以修改值,但是修改引用没用
func deferChanData() map[string]string {
	a := map[string]string{
		"name": "yuliyang",
	}
	defer func() {
		//a["name"] = "changed"
		a = map[string]string{
			"age": "22",
		}
	}()
	return a
}
