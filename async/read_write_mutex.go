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
		//写锁,不能重复加锁
		rw.Lock()
		//rw.Lock()
		fmt.Println("wlocking")
		time.Sleep(time.Second * 1)
		b++
		//rw.Unlock()
		rw.Unlock()
	}()
	time.Sleep(time.Millisecond * 200)
	go func() {
		//读锁加了几次锁，就要调用几次解锁
		rw.RLock()
		rw.RLock()
		rw.RLock()
		fmt.Println("rlocking")
		time.Sleep(time.Second * 1)
		fmt.Println(b)
		rw.RUnlock()
		rw.RUnlock()
		rw.RUnlock()
	}()
	time.Sleep(time.Millisecond * 5000)
	//写锁
	rw.Lock()
	fmt.Println("wlocking")
	time.Sleep(time.Second * 1)
	b++
	rw.Unlock()
	fmt.Println(b)
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
