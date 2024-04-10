package service

import (
	"github.com/go-playground/validator/v10"
	"library/data/request"
	"library/data/response"
	"library/helper"
	"library/model"
	"library/repository"
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
	helper.ErrorPanic(err)

	bookModel := model.Book{
		Name:   book.Name,
		Page:   book.Page,
		Author: book.Author,
		Status: book.Status,
	}

	b.BookRepository.Save(bookModel)
}

func (b BookServiceImpl) Update(book request.UpdateBookRequest) {
	bookData, err := b.BookRepository.FindById(book.Id)
	helper.ErrorPanic(err)
	bookData.Status = book.Status
	b.BookRepository.Update(*bookData)

}

func (b BookServiceImpl) Delete(bookId int) {
	b.BookRepository.Delete(bookId)
}

func (b BookServiceImpl) FindById(bookId int) *response.BooksResponse {
	book, err := b.BookRepository.FindById(bookId)
	helper.ErrorPanic(err)
	
	if book == nil {

		bookData := response.BooksResponse{
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
			Name:   value.Name,
			Author: value.Author,
			Status: value.Status,
		}
		books = append(books, book)
	}
	return books
}
