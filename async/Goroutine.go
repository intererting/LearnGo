package main

import (
	"fmt"
	"time"
)

var a = 0

func main() {
	go myGoroutine()
	fmt.Printf("after")
	myGoroutine()
}

func myGoroutine() {
	for i := 0; i < 10; i++ {
		a++
		time.Sleep(100 * time.Millisecond)
		fmt.Println(fmt.Sprintf("%d", a))
	}
}
