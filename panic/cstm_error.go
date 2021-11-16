package main

/**
自定义Error
系统自带error接口
type error interface {
    Error() string
}
*/
func testCstmError() error {
	a := myError{}
	return a
}

type myError struct {
}

//Error方法相当于toString，但是是error类型
func (m myError) Error() string {
	return "catch error"
}
