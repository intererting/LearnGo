package main

//闭包
func main() {
	functions := closureTest()
	println(functions(1))
	sum = 3
	println(functions(2))
}

var sum = 0

func closureTest() func(int) int {

	//闭包引用闭包外的值
	returnFunc := func(i int) int {
		return sum + i
	}
	//sum的值改变会影响闭包
	sum++
	return returnFunc
}
