package app

import (
	"encoding/json"
	"fmt"
	"library/cache"
	"library/data/request"
	"library/data/response"
	"library/helper"
	"library/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BookApp struct {
	bookService service.BookService
}

var (
	cacheRedis = cache.NewRedisClient("localhost:6379")
)

func NewBookApp(service service.BookService) *BookApp {
	return &BookApp{
		bookService: service,
	}
}

func (bookApp *BookApp) Create(ctx *fiber.Ctx) error {
	crateBookRequest := request.CreateBookRequest{}
	err := ctx.BodyParser(&crateBookRequest)

	helper.ThrowError(err)
	bookApp.bookService.Create(crateBookRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "OK",
		Message: "Successfully created a book",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)

}

func (bookApp *BookApp) Update(ctx *fiber.Ctx) error {
	updateBookRequest := request.UpdateBookRequest{}
	err := ctx.BodyParser(&updateBookRequest)
	helper.ThrowError(err)

	bookId := ctx.Params("bookId")
	id, err := strconv.Atoi(bookId)
	helper.ThrowError(err)

	updateBookRequest.Id = id
	bookApp.bookService.Update(updateBookRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "OK",
		Message: "Successfully updated a book",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)

}

func (bookApp *BookApp) Delete(ctx *fiber.Ctx) error {
	bookId := ctx.Params("bookId")
	id, err := strconv.Atoi(bookId)
	helper.ThrowError(err)
	bookApp.bookService.Delete(id)

	webResponse := response.Response{
		Code:    200,
		Status:  "OK",
		Message: "Successfully deleted a book",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)

}

func (bookApp *BookApp) FindById(ctx *fiber.Ctx) error {
	bookId := ctx.Params("bookId")
	id, err := strconv.Atoi(bookId)
	helper.ThrowError(err)
	bookResponse := bookApp.bookService.FindById(id)

	webResponse := response.Response{
		Code:    200,
		Status:  "OK",
		Message: "Successfully get a book",
		Data:    bookResponse,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)

}

func (bookApp *BookApp) FindAll(ctx *fiber.Ctx) error {

	var bookResponse []response.BooksResponse

	booksCache := cacheRedis.Get("books")

	if booksCache !=nil {
		
	
		err := json.Unmarshal(booksCache,&bookResponse)
		fmt.Println("Burada var")

		if err != nil {
			return nil
		}


	}else{

		bookResponse = bookApp.bookService.FindAll()

		booksListBytes , _ := json.Marshal(bookResponse)

		go func(books []byte) {
			fmt.Println("Kaydedildi")
			cacheRedis.Set("books",books)
		}(booksListBytes)
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "OK",
		Message: "Successfully get a book",
		Data:    bookResponse,
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)

}
