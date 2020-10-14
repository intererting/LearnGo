package main

import "fmt"

/**
安全的类型转换 as?
*/
func main() {
	//只有a为int类型才能转换成功
	var a interface{}
	//b, ok := a.(int)
	//fmt.Println(b, " ", ok)
}

/**
.(type)只能在switch中使用**
 */
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
