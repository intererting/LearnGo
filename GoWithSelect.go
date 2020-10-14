package main

import "fmt"

func fibonacciWithSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		//select 会阻塞到某个分支可以继续执行为止，这时就会执行该分支。当多个分支都准备好时会随机选择一个执行。
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		//goroutine等待通道中的数据
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		//输出了10个以后向quit发送结束命令
		quit <- 0
	}()
	fibonacciWithSelect(c, quit)
}
