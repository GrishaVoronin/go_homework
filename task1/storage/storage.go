package storage

import "lib/models"

type BookStorage interface {
	AddBook(b *models.Book, id string)
	Search(id string) (models.Book, bool)
}
