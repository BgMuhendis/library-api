package repository

import (
	"errors"
	"library/handlers"
	"library/model/dto/request"
	"library/model/entity"

	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
	Db *gorm.DB
}

func NewBookRepositoryImpl(Db *gorm.DB) BookRepository {
	return BookRepositoryImpl{Db: Db}
}

func (b BookRepositoryImpl) Save(book entity.Book) {
	result := b.Db.Create(&book)
	handlers.ThrowError(result.Error)
}

func (b BookRepositoryImpl) Update(book entity.Book) {
	var updateBook = request.UpdateBookRequest{
		Id:     book.Id,
		Status: book.Status,
	}
	result := b.Db.Model(&book).Where("status = ?",!updateBook.Status).Update("status",updateBook.Status)
	handlers.ThrowError(result.Error)

}

func (b BookRepositoryImpl) Delete(bookId int) {
	var book entity.Book
	result := b.Db.Where("id=?", bookId).Delete(&book)
	handlers.ThrowError(result.Error)

}

func (b BookRepositoryImpl) FindById(bookId int) (*entity.Book, error) {
	var book entity.Book
	result := b.Db.Find(&book, bookId)
	if result.RowsAffected != 1 {
		return nil, errors.New("Book not found")
	}
	return &book, nil
}

func (b BookRepositoryImpl) FindAll() []entity.Book {
	var books []entity.Book
	result := b.Db.Order("id").Find(&books)
	handlers.ThrowError(result.Error)
	return books
}
