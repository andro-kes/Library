package app

import (
	"Library/internal/database"
	"fmt"
	"log"
)

func Run(){
	database.CreateDataBase()
	getBook()
}

func getBook(){
	fmt.Println("Какую книгу вы ищете?")
	var id int
	fmt.Scanln(&id)
	book, ok := database.BooksBase.GetElementById(id)
	if ok {
		fmt.Println(book.GetTitle())
	} else {
		log.Println("Такой книги нет")
	}
}
