package main

import "fmt"

var a = 0

func testReference() {
	b := func() {
		a++
	}
	fmt.Println(a)
	b()
	fmt.Println(a)
}
