package main

import (
	"context"
	"fmt"
)

var myKey int

func contextWithValue() {
	ctx := context.WithValue(context.Background(), &myKey, "yuliyang")

	ctxInner, _ := context.WithCancel(ctx)
	fmt.Println(ctxInner.Value(&myKey))
}
