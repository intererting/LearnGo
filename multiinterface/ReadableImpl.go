package main

import "fmt"

type ReadableImpl struct {
	content string
}

func (r ReadableImpl) write(content string) string {
	return "write " + content
}

func (r ReadableImpl) read(content string) {
	fmt.Println(content)
}
