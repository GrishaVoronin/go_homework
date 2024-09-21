package main

type BookAdder interface {
	AddBook(book *Book, id string)
	ChangeStorage(*BookStorage)
}

type MyBookAdder struct {
	storage BookStorage
}

func NewMyBookAdder(storage *BookStorage) *MyBookAdder {
	return &MyBookAdder{*storage}
}

func (myAdder *MyBookAdder) AddBook(book *Book, id string) {
	myAdder.storage.AddBook(book, id)
}

func (myAdder *MyBookAdder) ChangeStorage(storage *BookStorage) {
	myAdder.storage = *storage
}
