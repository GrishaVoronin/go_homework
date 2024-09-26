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

func NewMyLibrary(s *storage.BookStorage, genId func(string) string) *MyLibrary {
	return &MyLibrary{*s, genId, make(map[string]string)}
}

func (myLib *MyLibrary) SetGeneratorId(genId func(string) string) {
	myLib.generateId = genId
}

func (m *MyLibrary) AddBook(book *models.Book) {
	id := m.generateId(book.Title)
	m.booksId[book.Title] = id
	m.Storage.AddBook(book, id)
}

func (m *MyLibrary) Search(name string) (models.Book, bool) {
	id, ok := m.booksId[name]
	if !ok {
		return models.Book{}, false
	}
	return m.Storage.Search(id)
}

func (m *MyLibrary) ChangeStorage(newStorage *storage.BookStorage) {
	m.Storage = *newStorage
	m.booksId = make(map[string]string)
}
