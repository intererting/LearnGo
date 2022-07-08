package main

import (
	"fmt"
	"reflect"
)

func testGoTag() {
	goTagS := GoTagS{
		name: "yu",
	}
	rType := reflect.TypeOf(goTagS)
	fmt.Println(rType.Field(0).Tag.Get("name"))
}

type GoTagS struct {
	name string `name:"yuliyang" age:"18"`
}
