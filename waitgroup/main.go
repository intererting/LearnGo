package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	go func() {
		time.Sleep(4 * time.Second)
		waitGroup.Add(-1)
	}()
	waitGroup.Add(1)
	waitGroup.Wait()
	log.Println("over")
}
