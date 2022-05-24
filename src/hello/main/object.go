package main

import "fmt"

type Book struct {
	name string
	auth string
}

func updateBook(book *Book) {
	book.auth = "222"
	fmt.Println("----- > %v", book)
}

//
//func main() {
//	var book Book
//	book.name = "name"
//	book.auth = "111"
//	fmt.Println("%v", book)
//
//	updateBook(&book)
//}
