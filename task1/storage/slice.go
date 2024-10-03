package storage

import "lib/models"

type SliceBookStorage struct {
	bookLocation map[string]int
	books        []models.Book
}

func NewSliceBookStorage() *SliceBookStorage {
	return &SliceBookStorage{make(map[string]int), []models.Book{}}
}

func (m *SliceBookStorage) AddBook(book *models.Book, id string) {
	m.bookLocation[id] = len(m.books)
	m.books = append(m.books, *book)
}

func (m *SliceBookStorage) Search(id string) (models.Book, bool) {
	ind := m.bookLocation[id]
	return m.books[ind], true
}
