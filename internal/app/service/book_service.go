package service

import (
	"library/internal/app/model/dto/request"
	"library/internal/app/model/dto/response"
)

type BookService interface {
	Create(book request.CreateBookRequest)
	Update(book request.UpdateBookRequest)
	Delete(bookId int)
	FindById(bookId int) *response.BooksResponse
	FindAll() []response.BooksResponse
}
