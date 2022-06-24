package main

import "fmt"

/**
安全的类型转换 as?
*/
func testTypeCast() {
	//只有a为int类型才能转换成功
	//var a interface{}
	//b, ok := a.(int)
	//fmt.Println(b, " ", ok)

	//type assertion,相当于as
	var b typeInterface
	b = typeStruct{}
	c := b.(typeInterface)
	//接口类型转换
	//d := typeInterface(b)
	c.content()
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

type typeInterface interface {
	content() string
}

type typeStruct struct {
}

func (t typeStruct) content() string {
	return "TypeStruct"
}

type typeStructTwo struct {
}

func (t typeStructTwo) content() string {
	return "TypeStructTwo"
}
