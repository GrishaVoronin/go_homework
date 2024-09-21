package main

import "fmt"

type Book struct {
	title       string
	author      string
	year        int
	pagesAmount int
	pages       []string
}

func (b *Book) print() {
	fmt.Printf("Title: %s\n", b.title)
	fmt.Printf("Author: %s\n", b.author)
	fmt.Printf("Year: %d\n", b.year)
	for _, page := range b.pages {
		fmt.Printf("%s\n", page)
	}
}

func (b *Book) getPage(page int) string {
	if page < 1 || page > b.pagesAmount {
		return "Not valid page number for this Book"
	}
	return b.pages[page-1]
}
