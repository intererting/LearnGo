package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func testByteBuffer() {
	file, fileErr := os.Open("test.txt")
	if fileErr != nil {
		return
	}
	defer file.Close()
	var data = make([]byte, 1024)

	var hasRead int
	var err error
	var byteBuffer bytes.Buffer

	hasRead, err = file.Read(data)
	for hasRead > 0 {
		//有数据的
		byteBuffer.Write(data[:hasRead])
		hasRead, err = file.Read(data)
		//捕获到了一个异常
		if err != nil && err != io.EOF {
			log.Fatal("error")
		}
	}
	fmt.Println(byteBuffer.String())
}
