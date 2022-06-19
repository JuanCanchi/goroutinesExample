package data

import (
	"fmt"
	"sync"
)

type Book struct {
	Id       int
	Title    string
	Finished bool
}

var books = []*Book{
	{1, "Dune", false},
	{2, "El Perfume", false},
	{3, "The World", false},
	{4, "Teoria de la noche", false},
	{5, "Ell principito", false},
	{6, "100 aos de soledad", false},
	{7, "El alquimista", false},
	{8, "El libro del cemeterio", false},
	{9, "Maze runner", false},
	{10, "Juan es GO", false},
}

func findBook(id int, m *sync.RWMutex) (int, *Book) {
	index := -1
	var book *Book

	m.RLock()
	for i, b := range books {
		if b.Id == id {
			index = i
			book = b
		}
	}
	m.RUnlock()
	return index, book
}

func FinishedBook(id int, m *sync.RWMutex) {
	i, book := findBook(id, m)
	if i < 0 {
		return
	}
	m.Lock()
	book.Finished = true
	books[i] = book
	m.Unlock()

	fmt.Printf("Finished book: %s\n", book.Title)
}
