package storage

import "lib/models"

type MapBookStorage struct {
	books map[string]models.Book
}

func NewMapBookStorage() *MapBookStorage {
	return &MapBookStorage{make(map[string]models.Book)}
}

func (m *MapBookStorage) AddBook(book *models.Book, id string) {
	m.books[id] = *book
}

func (m *MapBookStorage) Search(id string) (models.Book, bool) {
	return m.books[id], true
}
