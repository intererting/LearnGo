package main

import (
	"fmt"
	"time"
)

func ChannelForeach() {
	channel := make(chan int)
	go func() {
		//for range 也是阻塞的
		for data := range channel {
			fmt.Println(data)
		}
		fmt.Println("over")
	}()
	time.Sleep(time.Second)

}
