package database

import (
	"encoding/csv"
	"io"
	"os"
	"log"
)

type db = map[int]Record

type Record struct {
	Title string
	Author string
	Genre string
	Height string
	Publisher string
}

func CreateDataBase(){
	id := 0
	booksBase := make(db)
	reader := createCSVReader()

	_, err := reader.Read()
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
}

func createCSVReader() *csv.Reader {
	file, err := os.Open("books.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	return reader
}

func createRecord(record []string) Record {
	return Record{
		Title: record[0],
		Author: record[1],
		Genre: record[2],
		Height: record[3],
		Publisher: record[4],
	}
}