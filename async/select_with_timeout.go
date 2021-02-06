package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	timeOut := time.After(5 * time.Second)
	a := make(chan int)
	stop := make(chan int)
	wg.Add(1)
	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				a <- time.Now().Second()
				time.Sleep(time.Second)
			}
		}
	}()
	go func() {
		for {
			select {
			case <-timeOut:
				fmt.Println("timeout")
				stop <- 0
				wg.Done()
				break
			case data := <-a:
				fmt.Println(data)
			}
		}
	}()

	wg.Wait()

}
