package repository

import (
	"library/data/request"
	"library/helper"
	"library/model"

	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
	Db *gorm.DB
}

func NewBookRepositoryImpl(Db *gorm.DB) BookRepository {
	return BookRepositoryImpl{Db: Db}
}

func (b BookRepositoryImpl) Save(book model.Book) {
	result := b.Db.Create(&book)
	helper.ThrowError(result.Error)
}

func (b BookRepositoryImpl) Update(book model.Book) {
	var updateBook = request.UpdateBookRequest{
		Id:     book.Id,
		Status: book.Status,
	}
	result := b.Db.Model(&book).Where("status = ?",!updateBook.Status).Update("status",updateBook.Status)
	helper.ThrowError(result.Error)

}

func (b BookRepositoryImpl) Delete(bookId int) {
	var book model.Book
	result := b.Db.Where("id=?", bookId).Delete(&book)
	helper.ThrowError(result.Error)

}

func (b BookRepositoryImpl) FindById(bookId int) (*model.Book, error) {
	var book model.Book
	result := b.Db.Find(&book, bookId)
	if result.RowsAffected == 1 {
		return &book, nil

	} else {
		return nil, nil
	}
}

func (b BookRepositoryImpl) FindAll() []model.Book {
	var books []model.Book
	result := b.Db.Find(&books)
	helper.ThrowError(result.Error)
	return books
}
