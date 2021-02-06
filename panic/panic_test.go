package main

import "log"

/**
defer中也可以抛出panic
panic内部recover优先级高
内部捕获以后,不会影响外函数运行
*/
func main() {

	defer func() {
		if r := recover(); r != nil {
			log.Println("catch panic outer ", r)
		}
	}()

	testDeferPanic()

	log.Println("after catch panic")
}

func testDeferPanic() {

	defer func() {
		if r := recover(); r != nil {
			log.Println("catch panic inner", r)
		}
	}()

	log.Println("before")

	defer func() {
		panic("throw panic")
	}()

	log.Println("after")
}
