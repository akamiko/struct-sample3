package book

type Book struct {
	Title  string
	Author string
	Price  int
}

func NewBook(title string) *Book {
	return &Book{
		Title:  title,
		Author: "Unknown Author",
		Price:  100,
	}
}

func (b *Book) GetTitle() string {
	return b.Title
}
