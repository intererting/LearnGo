package main

import (
	"fmt"
	"reflect"
)

func testMuiltInterface() {
	var a readableWriteable
	a = ReadableImpl{"haha"}
	b := ReadableImpl{"haha"}
	fmt.Println(reflect.DeepEqual(a, b))
	fmt.Println(a == b)
	a.read("hello")
	fmt.Println(a.write("go"))
}

type readable interface {
	read(content string)
}

type writeable interface {
	write(content string) string
}

type readableWriteable interface {
	readable
	writeable
}

type ReadableImpl struct {
	content string
}

func (r ReadableImpl) write(content string) string {
	return "write " + content
}

func (r ReadableImpl) read(content string) {
	fmt.Println(content)
}
