package main

import (
	"fmt"
	"sync"
	"time"
)

/**
pool是协程安全的
pool只能被取一次，后面的会是nil
*/
func main() {
	var pool sync.Pool
	wg.Add(1)
	pool.Put("haha")
	pool.Put("next")
	go func() {
		fmt.Println(pool.Get())
		fmt.Println(pool.Get())
	}()
	go func() {
		fmt.Println(pool.Get())
	}()
	go func() {
		fmt.Println(pool.Get())
	}()

	go func() {
		_, ok := <-time.After(time.Second * 2)
		if ok {
			wg.Done()
		}
	}()
	wg.Wait()
}
