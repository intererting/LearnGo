package main

import "fmt"

func main() {
	var person = Person{
		price: 33.3,
		Book: Book{
			price: 22.2,
		},
	}
	//fmt.Println(person.price)
	//fmt.Println(person.Book.price)

	//person.test()

	fmt.Println(person)
}

type Person struct {
	price float64
	//匿名字段，匿名字段可以直接方法，当存在重命的时候外部优先
	//匿名字段里面的方法，也可以被直接访问
	Book
}

type Book struct {
	price float64
}

func (book *Book) test() {
	book.price = 300
}
