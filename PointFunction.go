package main

//函数和指针
func main() {
	myPoint := myPoint{
		x: 0,
		y: 0,
	}
	//扩展需要值类型的地方传指针没用,但是需要指针时可以直接用变量
	myPointP := &myPoint
	myPointP.test1()
	//test3(&myPoint)
	//fmt.Println(myPoint)
}

type myPoint struct {
	x float64
	y float64
}

//如果这里不带指针，那么不能修改值
func (myPoint myPoint) test1() {
	myPoint.x = 3
	myPoint.y = 4
}

//结构体也是值传递的，和swift一样，每次传递都是一次拷贝（如果用指针可以避免大数据的拷贝），所以这里不用指针不能修改值
func test2(point myPoint) {
	point.x = 3
	point.y = 4
}

func test3(point *myPoint) {
	point.x = 3
	point.y = 4
}
