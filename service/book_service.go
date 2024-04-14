package service

import (
	"library/data/request/book"
	"library/data/response/book"
)

type BookService interface {
	Create(book request.CreateBookRequest)
	Update(book request.UpdateBookRequest)
	Delete(bookId int)
	FindById(bookId int) *response.BooksResponse
	FindAll() []response.BooksResponse
}
