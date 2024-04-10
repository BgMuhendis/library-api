package repository

import (
	"errors"
	"gorm.io/gorm"
	"library/data/request"
	"library/helper"
	"library/model"
)

type BookRepositoryImpl struct {
	Db *gorm.DB
}

func NewBookRepositoryImpl(Db *gorm.DB) BookRepository {
	return BookRepositoryImpl{Db: Db}
}

func (b BookRepositoryImpl) Save(book model.Book) {
	result := b.Db.Create(&book)
	helper.ErrorPanic(result.Error)
}

func (b BookRepositoryImpl) Update(book model.Book) {
	var updateBook = request.UpdateBookRequest{
		Status: book.Status,
	}

	result := b.Db.Model(&book).Updates(updateBook)
	helper.ErrorPanic(result.Error)

}

func (b BookRepositoryImpl) Delete(bookId int) {
	var book model.Book
	result := b.Db.Where("id=?", bookId).Delete(&book)
	helper.ErrorPanic(result.Error)

}

func (b BookRepositoryImpl) FindById(bookId int) (*model.Book, error) {
	var book model.Book
	result := b.Db.Find(&book, bookId)
	if result != nil {
		return &book, nil
	} else {
		return nil, errors.New("Book is not found")
	}
}

func (b BookRepositoryImpl) FindAll() []model.Book {
	var books []model.Book
	result := b.Db.Find(&books)
	helper.ErrorPanic(result.Error)
	return books
}
