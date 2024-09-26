package library

import (
	"lib/models"
	"lib/storage"
)

type Library interface {
	Search(name string) (models.Book, bool)
	AddBook(*models.Book)
	SetGeneratorId(func(string) string)
	ChangeStorage(storage *storage.BookStorage)
}
