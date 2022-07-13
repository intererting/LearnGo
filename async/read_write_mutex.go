package main

import (
	"fmt"
	"sync"
	"time"
)

var rw sync.RWMutex
var b = 0

func rwmutexTest1() {
	go func() {
		rw.Lock()
		fmt.Println("wlocking")
		time.Sleep(time.Second * 3)
		b++
		rw.Unlock()
	}()
	time.Sleep(time.Millisecond * 200)
	go func() {
		rw.RLock()
		fmt.Println("rlocking")
		time.Sleep(time.Second * 3)
		fmt.Println(b)
		rw.RUnlock()
	}()

	time.Sleep(time.Second * 10)
}

func rwmutexTest2() {
	go func() {
		rw.RLock()
		fmt.Println("rlocking----1")
		time.Sleep(time.Second * 3)
		fmt.Println(b)
		rw.RUnlock()
	}()
	go func() {
		rw.RLock()
		fmt.Println("rlocking----2")
		time.Sleep(time.Second * 3)
		fmt.Println(b)
		rw.RUnlock()
	}()

	time.Sleep(time.Second * 10)
}
