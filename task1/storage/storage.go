package storage

import "lib/models"

type BookStorage interface {
	AddBook(book *models.Book, id string)
	Search(id string) (models.Book, bool)
}
