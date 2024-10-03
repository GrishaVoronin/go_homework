package library

import (
	"lib/models"
)

type Library interface {
	Search(name string) (models.Book, bool)
	AddBook(*models.Book)
	SetGeneratorId(func(string) string)
}
