package models

type BookBuilder struct {
	id       int
	bookName string
	price    float64
}

func (this *BookBuilder) SetPrice(price float64) *BookBuilder {
	this.price = price
	return this
}

func NewBookBuilder(id int, bookName string) *BookBuilder {
	return &BookBuilder{id: id, bookName: bookName}
}

func (this *BookBuilder) Build() *Book {
	book := &Book{Id: this.id, BookName: this.bookName}
	if this.price > 0 {
		book.Price = this.price
	}
	return book

}
