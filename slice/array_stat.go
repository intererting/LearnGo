package main

func learnArray() {

	////数组初始化
	//a := [12]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	////数组切片，前闭后开
	////切片后不再是数组
	//b := a[0:2]
	//fmt.Printf("%v\n", reflect.TypeOf(b))
	//////切片赋值会影响原数组
	//b[0] = 11
	//fmt.Println(a)
	//c := a[:]
	//c[0] = 100
	//fmt.Println(a)

	//cap在切换中表示下次切片能切的右值
	//a := make([]int, 0, 5)
	//fmt.Println(a)
	//b := a[1:2]
	//fmt.Println(b)
	//fmt.Println(len(b), cap(b))
	////这里只能切到4,切片引用的原数组
	//c := b[0:4]
	//fmt.Println(c)
	//fmt.Println(len(c), cap(c))

	//append追加
	//var a []int
	//a = append(a, 3)
	//fmt.Println(a)

	//range循环
	//注意点，range遍历获得的value是数据的副本，不是源数据的引用(指针，map，chan除外，这些是引用)
	//a := []int{'a', 'b'}
	//for index, value := range a {
	//	fmt.Println(index)
	//	//字符串可以用%q打印
	//	fmt.Printf("%q \n", value)
	//}

	//指针切片
	//a := 1
	//b := 2
	//c := []*int{&a, &b}
	//for _, value := range c {
	//	*value = 100
	//}
	//for _, value := range c {
	//	fmt.Println(*value)
	//}

	//相同长度，相同类型的数组可以直接赋值,这种赋值是拷贝
	//var a [3]int
	//b := [3]int{1, 2, 3}
	//a = b
	//a[0] = 100
	//fmt.Println(a)
	//fmt.Println(b)

	//a := []int{1, 2, 3, 4, 5, 6}
	//第一个数表示切片开始位置
	//第二个数表示len结束位置
	//第三个数表示cap结束位置
	//b := a[1:3:5]
	//fmt.Println(len(b), cap(b))
}
