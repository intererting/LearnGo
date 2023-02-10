package main

import "fmt"

//切片包含len和cap
//切片是对源数据的引用，所以修改切片会影响源数据
//如果切片的len<cap,那么想切片添加数据的时候，会直接修改源数据，相当于覆盖了元数据
//数据安全性：
//如果切片的len==cap，那么往切片添加数据的时候，会生成一个新的切片，这个切片和源数据没有关系
func main() {
	//testArray()
	//testChangeSlice()
	//testMap()
	//testChan()
	testSlice()
}

func testSlice() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := a[1:3:4]
	//最后一个4声明cap为4，表示该切片cap为a[4]，表示可以扩容的上界
	fmt.Println(len(b))
	fmt.Println(cap(b))
	c := append(b, 2)
	//超过了cap,那么就会生成一个新的切片
	d := append(c, 3)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func testSlice1() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := a[1:3]
	fmt.Println(len(b))
	fmt.Println(cap(b))
	c := append(b, 2)
	d := append(c, 3)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

//go的引用类型包括切片 map chan 函数 接口

//只要声明了长度就是数组，数组是拷贝，不是引用类型,所以不能修改
func testArray() {
	a := [2]int{1, 2}
	change(a)
	fmt.Print(a)
}

func change(a [2]int) {
	a[0] = 110
}

//这种情况是切片，切片是引用类型
func testChangeSlice() {
	a := []int{1, 2}
	changeSlice(a)
	fmt.Print(a)
}

func changeSlice(a []int) {
	a[0] = 110
}

//map也是引用类型
func testMap() {
	a := map[string]string{"name": "yuliyang"}
	changeMap(a)
	fmt.Println(a)
}

func changeMap(a map[string]string) {
	a["name"] = "changed"
}

//chan也是引用类型
func testChan() {
	a := make(chan string)
	go changeChan(a)
	b := <-a
	fmt.Println(b)
}

func changeChan(a chan<- string) {
	a <- "changed"
}
