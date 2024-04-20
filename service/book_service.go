package service

import (
	"library/model/dto/request"
	"library/model/dto/response"
)

type BookService interface {
	Create(book request.CreateBookRequest)
	Update(book request.UpdateBookRequest)
	Delete(bookId int)
	FindById(bookId int) *response.BooksResponse
	FindAll() []response.BooksResponse
}
