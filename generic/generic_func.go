package main

import (
	"fmt"
	"reflect"
)

func genericFunc[T int | string](t T) {
	fmt.Println(reflect.TypeOf(t))
}
