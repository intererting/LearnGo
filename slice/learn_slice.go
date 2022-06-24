package main

import "fmt"

func testSlice() {
	a := []int{1, 2}
	//fmt.Println(len(a), cap(a))
	sliceA := a[0:1]
	//sliceA的len为1，容量为可以再次扩容的容量，并且持有a的引用
	//fmt.Println(len(sliceA), cap(sliceA))
	appendResult := append(a, 1)
	//发生扩容，生成新的slice，这个时候的cap翻倍
	//fmt.Println(len(appendResult), cap(appendResult))
	//appendResult和之前的a没有关系了
	appendResult[0] = 1000
	fmt.Println(a)
	//尝试修改sliceA的值
	sliceA[0] = 100
	//a的值会被修改
	fmt.Println(a[0])
	//appendResult是新的slice,不会发生修改
	fmt.Println(appendResult[0])
}
