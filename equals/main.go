package main

import "fmt"

//测试go的==
func main() {
	//testSlice()

	//testArray()

	//testPointer()

	//testMap()

	//testChan()

	//testStruct()

	testNil()
}

func testNil() {
	type MyStruct struct {
	}

	type MyInterface interface {
	}

	var a []int
	var b MyStruct
	var c chan int
	var d map[string]string
	var e MyInterface
	var f *int
	var g func(int) int
	//引用类型默认为nil
	fmt.Println(a == nil)
	fmt.Println(c == nil)
	fmt.Println(d == nil)
	fmt.Println(e == nil)
	fmt.Println(f == nil)
	fmt.Println(g == nil)
	//基本类型会默认初始化
	fmt.Println(b)
}

//如果struct包含不能比较类型,比如切片,那么就会报错,其他情况只要类型可比较就行
func testStruct() {
	structA := MyStruct{
		//c: make(chan int),
	}
	structB := MyStruct{}
	fmt.Printf("%+v", structA)
	fmt.Println(structA == structB)
}

type MyStruct struct {
	//a []int
	a int
	//b map[string]string
	//c为nil
	c chan int
	d [3]int
}

//chan返回false
func testChan() {
	//chan默认nil
	var a chan int
	fmt.Printf("%v", a)
	chanA := make(chan int)
	chanB := make(chan int)
	fmt.Printf("%v", chanA)
	fmt.Println(chanA == chanB)
}

//map 不能比较map can only be compared to nil
func testMap() {
	//mapA := map[string]string{
	//	"name": "yuliyang",
	//}
	//mapB := map[string]string{
	//	"name": "yuliyang",
	//}
	//fmt.Println(mapA == mapB)
}

// 3指针,指针在指向同一个地址的时候是相等的
func testPointer() {
	//pointA := new(int)
	//pointB := new(int)
	//fmt.Println(pointA == pointB)

	a := 3
	pointA := &a
	pointB := &a
	fmt.Println(pointB == pointA)

}

//	2数组 数组大小不一样不能比较  数组只有内容都相同才能比较
func testArray() {
	arrayA := [3]int{1, 2, 3}
	arrayB := [3]int{1, 2, 3}
	fmt.Println(arrayA == arrayB)
}

//1:切片 slice can only be compared to nil
//切片不能直接比较,但是可以和nil比较
func testSlice() {
	//sliceA := []int{1, 2, 3}
	//sliceB := []int{1, 2, 3}
	////error
	//fmt.Println(sliceA == sliceB)
	//fmt.Println(nil == sliceA)
}
