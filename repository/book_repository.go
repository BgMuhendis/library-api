package repository

import "library/model"

type BookRepository interface {
	Save(book model.Book)
	Update(book model.Book)
	Delete(bookId int)
	FindById(bookId int) (*model.Book, error)
	FindAll() []model.Book
}
