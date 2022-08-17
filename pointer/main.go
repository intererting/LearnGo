package main

import "fmt"

func main() {
	arr := []int{1, 2}
	var a *[]int = &arr
	fmt.Println(*a)
}
