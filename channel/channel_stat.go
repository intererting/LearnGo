package main

import "fmt"

func testChannel() {

	//a := []int{1, 2, 3, 4}
	////2代表缓冲区大小
	//channel := make(chan int, 2)
	//go channelTest(a[:2], channel)
	//go channelTest(a[2:], channel)
	////祖塞，而且是栈，和defer差不多
	//left, right := <-channel, <-channel
	//fmt.Println(left, right)

	//测试斐波那契数列

	channel := make(chan int, 10)
	fibonacci(cap(channel), channel)
	//for value := range channel {
	//	fmt.Println(value)
	//}

	for {
		value, ok := <-channel
		if ok {
			fmt.Println(value)
		} else {
			break
		}
	}
}

/**
声明一个int类型的通道
*/
func channelTest(a []int, b chan<- int) {
	var sum = 0
	for _, value := range a {
		sum += value
	}
	b <- sum
}

/**
通道实现斐波那契数列
*/
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		//写入通道
		c <- x
		x, y = y, x+y
	}
	//手动关闭
	close(c)
}
