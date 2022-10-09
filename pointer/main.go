package main

import (
	"fmt"
	"reflect"
)

func main() {
	//arr := []int{1, 2}
	//var a *[]int = &arr
	//fmt.Println(*a)

	a := new(*int)
	fmt.Println(reflect.TypeOf(a))

}
