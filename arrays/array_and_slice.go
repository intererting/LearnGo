package main

import "fmt"

//go的引用类型包括map chan 切片 函数 接口

//这种情况是数组，数组是拷贝，不是引用类型
//func main() {
//	a := [2]int{1, 2}
//	change(a)
//	fmt.Print(a)
//}
//
//func change(a [2]int) {
//	a[0] = 110
//}

//这种情况是切片，切片是引用类型
//func main() {
//	a := []int{1, 2}
//	change(a)
//	fmt.Print(a)
//}
//
//func change(a []int) {
//	a[0] = 110
//}

//切片包含len和cap
//切片是对源数据的引用，所以修改切片会影响源数据
//如果切片的len<cap,那么想切片添加数据的时候，会直接修改源数据，相当于覆盖了元数据
//数据安全性：
//如果切片的len==cap，那么往切片添加数据的时候，会生成一个新的切片，这个切片和源数据没有关系
func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := a[:3:3]
	c := append(b, 2)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

//map也是引用类型
//func main() {
//	a := map[string]string{"name": "yuliyang"}
//	change(a)
//	fmt.Println(a)
//}
//
//func change(a map[string]string) {
//	a["name"] = "changed"
//}

//chan也是引用类型
//func main() {
//	a := make(chan string)
//	go change(a)
//	b := <-a
//	fmt.Println(b)
//}
//
//func change(a chan string) {
//	a <- "changed"
//}
