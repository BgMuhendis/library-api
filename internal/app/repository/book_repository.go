package repository

import "library/internal/app/model/entity"



type BookRepository interface {
	Save(book entity.Book)
	Update(book entity.Book)
	Delete(bookId int)
	FindById(bookId int) (*entity.Book, error)
	FindAll() []entity.Book
}
