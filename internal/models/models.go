package models

import (
	"Library/internal/book"
	"log"
)

type DB map[int]book.Book

func (db *DB) GetElementById(id int) (book.Book, bool) {
	book, ok := (*db)[id]
	return book, ok
}