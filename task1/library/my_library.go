package library

import (
	"lib/models"
	"lib/storage"
)

type MyLibrary struct {
	Storage    storage.BookStorage
	generateId func(string) string
	booksId    map[string]string
}

func NewMyLibrary(storage *storage.BookStorage, genId func(string) string) *MyLibrary {
	return &MyLibrary{*storage, genId, make(map[string]string)}
}

func (myLib *MyLibrary) SetGeneratorId(genId func(string) string) {
	myLib.generateId = genId
}

func (myLib *MyLibrary) AddBook(book *models.Book) {
	id := myLib.generateId(book.Title)
	myLib.booksId[book.Title] = id
	myLib.Storage.AddBook(book, id)
}

func (myLib *MyLibrary) Search(name string) (models.Book, bool) {
	id, ok := myLib.booksId[name]
	if !ok {
		return models.Book{}, false
	}
	return myLib.Storage.Search(id)
}

func (myLib *MyLibrary) ChangeStorage(newStorage *storage.BookStorage) {
	myLib.Storage = *newStorage
	myLib.booksId = make(map[string]string)
}
