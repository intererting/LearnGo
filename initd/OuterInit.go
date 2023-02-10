package main

import "fmt"

func init() {
	fmt.Println("Outer init 1")
}

func init() {
	fmt.Println("Outer init 2")
}
