package main

type Searcher interface {
	Search(id string) (Book, bool)
	ChangeStorage(*BookStorage)
}

type MySearcher struct {
	storage BookStorage
}

func NewMySearcher(storage *BookStorage) *MySearcher {
	return &MySearcher{*storage}
}

func (mySearcher *MySearcher) Search(id string) (Book, bool) {
	return mySearcher.storage.Search(id)
}

func (mySearcher *MySearcher) ChangeStorage(storage *BookStorage) {
	mySearcher.storage = *storage
}
