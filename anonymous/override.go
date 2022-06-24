package main

import "fmt"

func testEmbedProperty() {
	var person = Person{
		price: 33.3,
		Book: Book{
			price: 22.2,
		},
	}
	//访问权限提升,不需要使用Book调用
	person.test()
	fmt.Printf("%+v", person)
}

type Person struct {
	price float64
	//匿名字段，匿名字段可以直接访问，当存在重命的时候外部优先
	//匿名字段里面的方法，也可以被直接访问
	//方法和属性访问权限都可以被提升
	Book
}

type Book struct {
	price float64
}

func (book *Book) test() {
	book.price = 300
}
