package main

type Library interface {
	Search(name string) (Book, bool)
	AddBook(*Book)
	ChangeStorage(storage *BookStorage)
	SetGeneratorId(func(string) string)
}

type MyLibrary struct {
	searcher   Searcher
	bookAdder  BookAdder
	generateId func(string) string
	booksId    map[string]string
}

func NewMyLibrary(searcher *Searcher, adder *BookAdder, genId func(string) string) *MyLibrary {
	return &MyLibrary{*searcher, *adder, genId, make(map[string]string)}
}

func (myLib *MyLibrary) SetGeneratorId(genId func(string) string) {
	myLib.generateId = genId
}

func (myLib *MyLibrary) AddBook(book *Book) {
	id := myLib.generateId(book.title)
	myLib.booksId[book.title] = id
	myLib.bookAdder.AddBook(book, id)
}

func (myLib *MyLibrary) Search(name string) (Book, bool) {
	id, ok := myLib.booksId[name]
	if !ok {
		return Book{}, false
	}
	return myLib.searcher.Search(id)
}

func (myLib *MyLibrary) ChangeStorage(newStorage *BookStorage) {
	myLib.searcher.ChangeStorage(newStorage)
	myLib.bookAdder.ChangeStorage(newStorage)
	myLib.booksId = make(map[string]string)
}
