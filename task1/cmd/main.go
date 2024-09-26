package main

import (
	"fmt"
	"lib/id_generators"
	"lib/library"
	"lib/models"
	"lib/storage"
)

func main() {
	//create books
	Pages := []string{"page1", "page2", "page3", "page4", "page5"}
	book1 := models.Book{Title: "Ночной дозор", Author: "Сергей Лукьяненко", Year: 1998, PagesAmount: 5, Pages: Pages}
	book2 := models.Book{Title: "Дневной дозор", Author: "Сергей Лукьяненко", Year: 2000, PagesAmount: 5, Pages: Pages}
	book3 := models.Book{Title: "Сумеречный дозор", Author: "Сергей Лукьяненко", Year: 2004, PagesAmount: 5, Pages: Pages}
	books := []models.Book{book1, book2, book3}
	//create library
	var cur_storage storage.BookStorage = storage.NewSliceBookStorage()
	var lib library.Library = library.NewMyLibrary(&cur_storage, id_generators.GenerateIdMD5)
	//add books
	for _, book := range books {
		lib.AddBook(&book)
	}
	//search books
	book, ok := lib.Search(book1.Title)
	if !ok {
		fmt.Println("Book not found")
	} else {
		book.Print()
	}
	book, ok = lib.Search(book2.Title)
	if !ok {
		fmt.Println("Book not found")
	} else {
		book.Print()
	}
	//change generator id
	lib.SetGeneratorId(id_generators.GenerateIdSHA256)
	book, ok = lib.Search(book3.Title)
	if !ok {
		fmt.Println("Book not found")
	} else {
		book.Print()
	}
	//change storage
	var newStorage storage.BookStorage = storage.NewSliceBookStorage()
	lib.ChangeStorage(&newStorage)
	//create books
	book4 := models.Book{Title: "Последний дозор", Author: "Сергей Лукьяненко", Year: 2006, PagesAmount: 5, Pages: Pages}
	book5 := models.Book{Title: "Новый дозор", Author: "Сергей Лукьяненко", Year: 2012, PagesAmount: 5, Pages: Pages}
	book6 := models.Book{Title: "Шестой дозор", Author: "Сергей Лукьяненко", Year: 2014, PagesAmount: 5, Pages: Pages}
	books = []models.Book{book4, book5, book6}
	//add books
	for _, book := range books {
		lib.AddBook(&book)
	}
	//search book
	book, ok = lib.Search(book5.Title)
	if !ok {
		fmt.Println("Book not found")
	} else {
		book.Print()
	}
}
