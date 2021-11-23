package main

import (
	"context"
	"fmt"
	"time"
)

func contextWithCancel() {
	ctx, cancelF := context.WithCancel(context.Background())
	ctxInner, cancelFInner := context.WithCancel(ctx)
	defer cancelFInner()
	// consumer
	go func(ctx context.Context) {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-ctx.Done():
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Println("ticker")
			}
		}
	}(ctxInner)
	time.Sleep(time.Second * 3)
	cancelF()
	fmt.Println("ctx canceled")
	time.Sleep(time.Second * 3)
	fmt.Println("background canceled")
}
