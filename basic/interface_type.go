package main

import "fmt"

//interface{}类型可以用过.(类型)进行强制类型转换
func interfaceCast() {
	myStruct := MyStruct{
		name: "yuliyang",
	}
	var a interface{} = myStruct
	result := a.(MyInterfacer)
	fmt.Println(result)
}

type MyStruct struct {
	name string
}

type MyInterfacer interface {
	test()
}

func (myStruct MyStruct) test() {
	fmt.Println("MyStruct test")
}
