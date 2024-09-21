package main

import "fmt"

func main() {
	//create books
	pages := []string{"page1", "page2", "page3", "page4", "page5"}
	book1 := Book{title: "Ночной дозор", author: "Сергей Лукьяненко", year: 1998, pagesAmount: 5, pages: pages}
	book2 := Book{title: "Дневной дозор", author: "Сергей Лукьяненко", year: 2000, pagesAmount: 5, pages: pages}
	book3 := Book{title: "Сумеречный дозор", author: "Сергей Лукьяненко", year: 2004, pagesAmount: 5, pages: pages}
	books := []Book{book1, book2, book3}
	//create library
	var storage BookStorage = NewMyBookStorage()
	var searcher Searcher = NewMySearcher(&storage)
	var adder BookAdder = NewMyBookAdder(&storage)
	var lib Library = NewMyLibrary(&searcher, &adder, GenerateId1)
	//add books
	for _, book := range books {
		lib.AddBook(&book)
	}
	//search books
	book, ok := lib.Search(book1.title)
	if !ok {
		fmt.Println("Book not found")
	} else {
		book.print()
	}
	book, ok = lib.Search(book2.title)
	if !ok {
		fmt.Println("Book not found")
	} else {
		book.print()
	}
	//change generator id
	lib.SetGeneratorId(GenerateId2)
	book, ok = lib.Search(book3.title)
	if !ok {
		fmt.Println("Book not found")
	} else {
		book.print()
	}
	//change storage
	lib.ChangeStorage()
	//create books
	book4 := Book{title: "Последний дозор", author: "Сергей Лукьяненко", year: 2006, pagesAmount: 5, pages: pages}
	book5 := Book{title: "Новый дозор", author: "Сергей Лукьяненко", year: 2012, pagesAmount: 5, pages: pages}
	book6 := Book{title: "Шестой дозор", author: "Сергей Лукьяненко", year: 2014, pagesAmount: 5, pages: pages}
	books = []Book{book4, book5, book6}
	//add books
	for _, book := range books {
		lib.AddBook(&book)
	}
	//search book
	book, ok = lib.Search(book5.title)
	if !ok {
		fmt.Println("Book not found")
	} else {
		book.print()
	}
}
