package main

import (
	"fmt"
	"sync"
	"time"
)

/**
pool是协程安全的
pool只能被取一次，后面的会是nil
pool是弱应用，不会阻止回收
*/
func testPool() {
	var wg sync.WaitGroup
	var pool sync.Pool
	wg.Add(1)
	//队列 FIFO
	pool.Put("haha")
	pool.Put("next")

	go func() {
		fmt.Println(pool.Get())
		fmt.Println(pool.Get())
	}()
	go func() {
		fmt.Println(pool.Get())
		pool.Put("new")
	}()
	go func() {
		fmt.Println(pool.Get())
	}()

	go func() {
		_, ok := <-time.After(time.Second * 2)
		//当有结果是为true
		fmt.Println(ok)
		if ok {
			wg.Done()
		}
	}()
	wg.Wait()
}
