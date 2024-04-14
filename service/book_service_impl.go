package service

import (
	"library/data/request"
	"library/data/response"
	"library/helper"
	"library/models"
	"library/repository"

	"github.com/go-playground/validator/v10"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
	validate       *validator.Validate
}

func NewBookServiceImpl(bookRepository repository.BookRepository, validate *validator.Validate) BookService {
	return BookServiceImpl{BookRepository: bookRepository, validate: validate}
}

func (b BookServiceImpl) Create(book request.CreateBookRequest) {
	err := b.validate.Struct(book)
	helper.ThrowError(err)

	bookModel := models.Book{
		Name:   book.Name,
		Page:   book.Page,
		Author: book.Author,
		Status: book.Status,
	}

	b.BookRepository.Save(bookModel)
}

func (b BookServiceImpl) Update(book request.UpdateBookRequest) {
	bookData, err := b.BookRepository.FindById(book.Id)
	helper.ThrowError(err)
	bookData.Status = book.Status
	b.BookRepository.Update(*bookData)

}

func (b BookServiceImpl) Delete(bookId int) {
	b.BookRepository.Delete(bookId)
}

func (b BookServiceImpl) FindById(bookId int) *response.BooksResponse {
	book, err := b.BookRepository.FindById(bookId)
	helper.ThrowError(err)
	
	if book != nil {

		bookData := response.BooksResponse{
			Id: book.Id,
			Name:   book.Name,
			Author: book.Author,
			Status: book.Status,
		}
		return &bookData
	}
	return nil

}

func (b BookServiceImpl) FindAll() []response.BooksResponse {
	result := b.BookRepository.FindAll()
	var books []response.BooksResponse

	for _, value := range result {
		book := response.BooksResponse{
			Id: value.Id,
			Name:   value.Name,
			Author: value.Author,
			Status: value.Status,
		}
		books = append(books, book)
	}
	return books
}
