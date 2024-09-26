package storage

import "lib/models"

type SliceBookStorage struct {
	bookLocation map[string]int
	books        []models.Book
}

func NewSliceBookStorage() *SliceBookStorage {
	return &SliceBookStorage{make(map[string]int), []models.Book{}}
}

func (myStorage *SliceBookStorage) AddBook(book *models.Book, id string) {
	myStorage.bookLocation[id] = len(myStorage.books)
	myStorage.books = append(myStorage.books, *book)
}

func (myStorage *SliceBookStorage) Search(id string) (models.Book, bool) {
	ind := myStorage.bookLocation[id]
	return myStorage.books[ind], true
}
