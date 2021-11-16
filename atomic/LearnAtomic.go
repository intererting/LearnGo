package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var a int32 = 1
	atomic.AddInt32(&a, 1)
	fmt.Println(a)
}
