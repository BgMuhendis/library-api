package app

import (
	"github.com/gofiber/fiber/v2"
	"library/data/request"
	"library/data/response"
	"library/helper"
	"library/service"
	"strconv"
)

type BookApp struct {
	bookService service.BookService
}

func NewBookApp(service service.BookService) *BookApp {
	return &BookApp{
		bookService: service,
	}
}

func (bookApp *BookApp) Create(ctx *fiber.Ctx) error {
	crateBookRequest := request.CreateBookRequest{}
	err := ctx.BodyParser(&crateBookRequest)

	helper.ErrorPanic(err)
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
	helper.ErrorPanic(err)

	bookId := ctx.Params("bookId")
	id, err := strconv.Atoi(bookId)
	helper.ErrorPanic(err)

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
	helper.ErrorPanic(err)
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
	helper.ErrorPanic(err)
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
	bookResponse := bookApp.bookService.FindAll()
	webResponse := response.Response{
		Code:    200,
		Status:  "OK",
		Message: "Successfully get a book",
		Data:    bookResponse,
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)

}
