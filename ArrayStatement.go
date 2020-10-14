package main

import "fmt"

func main() {

	////数组初始化
	//a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	////数组切片，前闭后开
	//b := a[0:2]
	////切片赋值会影响原数组
	//b[0] = 11
	////fmt.Printf("%v", a[0:2])
	//fmt.Println(a[:])

	//cap在切换中表示下次切片能切的右值
	//a := make([]int, 0, 5)
	//fmt.Println(a)
	//b := a[1:2]
	//fmt.Println(b)
	//fmt.Println(len(b), cap(b))
	////这里只能切到4
	//c := b[0:4]
	//fmt.Println(c)
	//fmt.Println(len(c), cap(c))

	//a := make([]int, 5)
	//b := a[0:2]
	//b[1] = 3
	//fmt.Println(b)
	//fmt.Println(a)

	//append追加
	//var a []int
	//a = append(a, 3)
	//fmt.Println(a)

	//range循环
	a := []int{'a', 'b'}
	for index, value := range a {
		fmt.Println(index)
		fmt.Printf("%q \n", value)
	}
}
