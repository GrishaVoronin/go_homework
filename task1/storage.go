package main

type BookStorage interface {
	AddBook(book *Book, id string)
	Search(id string) (Book, bool)
}

type SliceBookStorage struct {
	bookLocation map[string]int
	books        []Book
}

func NewSliceBookStorage() *SliceBookStorage {
	return &SliceBookStorage{make(map[string]int), []Book{}}
}

func (myStorage *SliceBookStorage) AddBook(book *Book, id string) {
	myStorage.bookLocation[id] = len(myStorage.books)
	myStorage.books = append(myStorage.books, *book)
}

func (myStorage *SliceBookStorage) Search(id string) (Book, bool) {
	ind := myStorage.bookLocation[id]
	return myStorage.books[ind], true
}

type MapBookStorage struct {
	books map[string]Book
}

func NewMapBookStorage() *MapBookStorage {
	return &MapBookStorage{make(map[string]Book)}
}

func (myStorage *MapBookStorage) AddBook(book *Book, id string) {
	myStorage.books[id] = *book
}

func (myStorage *MapBookStorage) Search(id string) (Book, bool) {
	return myStorage.books[id], true
}
