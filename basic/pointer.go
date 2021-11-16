package main

func testPointer() {
	////指针声明
	//var a *int
	//var b = 3
	////指针赋值
	//c := &b
	//a = &b
	////指针修改值
	//*c = 4
	////指针获取值
	//d := *c
	//fmt.Printf("%p,%p,%d", a, c, d)
	//没有指针运算

	//go不像C,数组名不是一个指针常量,不能当做第一个元素的地址
	//var a = []int{1, 2, 3}
	//testArrayPointer(a)
	//log.Print(a)

}

func testArrayPointer(a *int) {
	*a = 100
}
