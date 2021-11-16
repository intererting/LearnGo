package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ioReader() {
	//基本的IO操作
	a := "hello golang"
	reader := strings.NewReader(a)
	b := make([]byte, 8)
	for {
		hasRead, err := reader.Read(b)
		if err == io.EOF {
			break
		}
		fmt.Printf("%s", b[:hasRead])
	}
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
