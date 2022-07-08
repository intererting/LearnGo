package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func testReflect() {
	//获取类型
	mStruct := myStruct{10, "yuliyang"}
	mType := reflect.TypeOf(mStruct)
	mName := mType.Field(1).Name
	fmt.Println(mName)

	//获取值
	mValue := reflect.ValueOf(mStruct).Field(1).String()
	fmt.Println(mValue)
}

//测试Pointer和uintptr
//uintptr表示一个地址常量，可以用来计算地址
//Pointer表示一个指针，可以通过指针获取值
func testPointer() {

	//通过地址的计算，获取当前地址的数据
	mStruct := myStruct{30, "yuliyang"}
	//计算offset
	namePointer := unsafe.Pointer(uintptr(unsafe.Pointer(&mStruct)) + unsafe.Offsetof(mStruct.name))
	convert := *(*string)(namePointer)
	fmt.Printf("%+v", convert)

	//fmt.Printf("%+v", *(*myStruct)(unsafe.Pointer(&mStruct)))
}

type myStruct struct {
	age  int32
	name string
}
