package main

import (
	"fmt"
)

func main() {
	//n := Node[int]{
	//	value: 5,
	//}
	n := Node_1[string]{
		value: "yuliyang",
	}
	fmt.Printf("%+v", n)
}

type Node[T any] struct {
	value T
}

type Node_1[T MyType] struct {
	value T
}

type MyType interface {
	int | string
}
