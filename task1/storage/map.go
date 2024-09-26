package storage

import "lib/models"

type MapBookStorage struct {
	books map[string]models.Book
}

func NewMapBookStorage() *MapBookStorage {
	return &MapBookStorage{make(map[string]models.Book)}
}

func (myStorage *MapBookStorage) AddBook(book *models.Book, id string) {
	myStorage.books[id] = *book
}

func (myStorage *MapBookStorage) Search(id string) (models.Book, bool) {
	return myStorage.books[id], true
}
