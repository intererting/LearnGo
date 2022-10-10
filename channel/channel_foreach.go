package main

import (
	"fmt"
	"time"
)

func ChannelForeach() {
	channel := make(chan int)
	go func() {
		for data := range channel {
			fmt.Println(data)
		}
		fmt.Println("over")
	}()
	time.Sleep(time.Minute)

}
