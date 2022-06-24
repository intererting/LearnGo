package main

import (
	"context"
	"fmt"
)

var myKey string

func contextWithValue() {
	ctx := context.WithValue(context.Background(), &myKey, "yuliyang")

	ctxInner, cancelF := context.WithCancel(ctx)
	defer cancelF()
	fmt.Println(ctxInner.Value(&myKey))
}
