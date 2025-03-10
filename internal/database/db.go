package database

import (
	"Library/internal/path"
	"Library/internal/book"
	"Library/internal/models"
	"encoding/csv"
	"io"
	"log"
	"os"
)



var BooksBase models.DB

func CreateDataBase() {
	id := 0
	booksBase := make(models.DB)

	file, err := os.Open(path.ABSOLUTE_PATH_TO_CSV)
	if err != nil {
		log.Fatal(err)
	}
	reader := createCSVReader(file)

	_, err = reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	for {
		record, err := reader.Read()

		if err == io.EOF {
			log.Println("База построена")
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		
		booksBase[id] = createRecord(record)
		id++
	}
	defer file.Close()
	BooksBase = booksBase
}

func createCSVReader(file *os.File) *csv.Reader {
	reader := csv.NewReader(file)
	return reader
}

func createRecord(record []string) book.Book {
	return book.Book{
		Title: record[0],
		Author: record[1],
		Genre: record[2],
		Height: record[3],
		Publisher: record[4],
	}
}