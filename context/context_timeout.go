package main

import (
	"context"
	"fmt"
	"time"
)

func contextTimeout() {
	messages := make(chan int, 10)
	defer close(messages)
	// producer
	for i := 0; i < 10; i++ {
		messages <- i
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// consumer
	go func(ctx context.Context) {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-ctx.Done():
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Printf("send message: %d\n", <-messages)
			}
		}
	}(ctx)

	//select {
	//case <-ctx.Done():
	//	fmt.Println("main.go process exit!")
	//}

	time.Sleep(time.Minute)
}
