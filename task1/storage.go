package main

type BookStorage interface {
	AddBook(book *Book, id string)
	Search(id string) (Book, bool)
}

type MyBookStorage struct {
	bookLocation map[string]int
	books        []Book
}

func NewMyBookStorage() *MyBookStorage {
	return &MyBookStorage{make(map[string]int), []Book{}}
}

func (myStorage *MyBookStorage) AddBook(book *Book, id string) {
	myStorage.bookLocation[id] = len(myStorage.books)
	myStorage.books = append(myStorage.books, *book)
}

func (myStorage *MyBookStorage) Search(id string) (Book, bool) {
	ind := myStorage.bookLocation[id]
	return myStorage.books[ind], true
}
