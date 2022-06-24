package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var a int32 = 0
var wg sync.WaitGroup

func testCoroutine() {
	wg.Add(2)
	go myGoroutine()
	go myGoroutine()
	wg.Wait()
	fmt.Println(fmt.Sprintf("%d", a))
}

func myGoroutine() {

	//多线程访问，可以用atomic修改
	for i := 0; i < 100000; i++ {
		atomic.AddInt32(&a, 1)
		//time.Sleep(100 * time.Millisecond)
	}
	defer wg.Done()
}

//当主进程结束，协程也被关闭
func goroutineDamon() {
	go func() {
		for {
			fmt.Println("in loop")
			time.Sleep(time.Second)
		}
	}()
}
