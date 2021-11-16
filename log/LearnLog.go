package main

import (
	"io"
	"log"
	"os"
)

func main() {
	//log.SetPrefix("haha prefix")
	//log.SetFlags(log.Ltime | log.LUTC)
	//log.Println("log")

	//自定义日志输出
	logFile, _ := os.OpenFile("/home/yuliyang/Desktop/test.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	newLog := log.New(io.Writer(logFile), "prefix", log.Ldate)
	newLog.SetFlags(log.Llongfile)
	newLog.Println("HAHA")
}
