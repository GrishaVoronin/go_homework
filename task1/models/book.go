package models

import "fmt"

type Book struct {
	Title       string
	Author      string
	Year        int
	PagesAmount int
	Pages       []string
}

func (b *Book) Print() {
	fmt.Printf("Title: %s\n", b.Title)
	fmt.Printf("Author: %s\n", b.Author)
	fmt.Printf("Year: %d\n", b.Year)
	for _, page := range b.Pages {
		fmt.Printf("%s\n", page)
	}
}

func (b *Book) GetPage(page int) string {
	if page < 1 || page > b.PagesAmount {
		return "Not valid page number for this Book"
	}
	return b.Pages[page-1]
}
