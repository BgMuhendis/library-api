package repository

import "library/models"

type BookRepository interface {
	Save(book models.Book)
	Update(book models.Book)
	Delete(bookId int)
	FindById(bookId int) (*models.Book, error)
	FindAll() []models.Book
}
