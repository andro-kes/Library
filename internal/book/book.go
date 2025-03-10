package book

type Book struct {
	Title string
	Author string
	Genre string
	Height string
	Publisher string
}

func (book *Book) getTitle() string {
	return book.Title
}