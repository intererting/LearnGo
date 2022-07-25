package main

/**
1:编译源代码 go build -gcflags "-N -l" main.go

2:gdb main

run 运行代码
list 查看源代码
break 设置断点
delete删除断点
backtrace 查看代码执行
info 查看变量
print 打印变量
whatis 查看类型
next 下一步
continue 跳过
set 设置变量
*/
func main() {
	var obj interface{}
	type User struct {
	}
	var u *User
	obj = u
	println(u == nil)   // true
	println(obj == nil) // false
}
