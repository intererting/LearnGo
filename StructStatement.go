package main

import "fmt"

func main() {

	type Person struct {
		name string
		age  int
	}

	person := Person{"yuliyang", 22}
	//带Name的命名
	personB := Person{name: "yuliyang2", age: 33}
	person.age = 30
	//结构体指针
	personP := &person
	//指针可以直接访问属性
	fmt.Println(personP.name)
	fmt.Printf("%v,%v", person, personB)
}
