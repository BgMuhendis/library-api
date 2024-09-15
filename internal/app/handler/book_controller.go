package controller

import (
	"encoding/json"
	"fmt"
	"library/cache"
	"library/errorHandler"
	"library/internal/app/model/dto/request"
	"library/internal/app/model/dto/response"
	"library/internal/app/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BookApp struct {
	bookService service.BookService
}

var (
	cacheRedis = cache.NewRedisClient("redis-cache:6379")
)


func NewBookApp(service service.BookService) *BookApp {
	return &BookApp{
		bookService: service,
	}
}

func cacheDelete() {

	cacheRedis.Del("books")
}

// CreateBook Create Order by Code
//	@Summary		Create Book by Code
//	@Description	Create Book by Code in detail
//	@Tags			Books
//	@Accept			json
//	@Produce		json
//	@Param			request	body	request.CreateBookRequest	true	"Request of Creating Book Object"
//	@Success		200		{json}	json
//	@Router			/api/book [post]
func (bookApp *BookApp) Create(ctx *fiber.Ctx) error {
	
	crateBookRequest := request.CreateBookRequest{}
	err := ctx.BodyParser(&crateBookRequest)

	handlers.ThrowError(err)
	bookApp.bookService.Create(crateBookRequest)

	webResponse := response.Response{
		Message: "Successfully created a book",
	}
	go cacheDelete()

	return ctx.Status(fiber.StatusOK).JSON(webResponse)

}

// UpdateBook Update Order by Code
//	@Summary		Update Book by Code
//	@Description	Update Book by Code in detail
//	@Tags			Books
//	@Accept			json
//	@Produce		json
//	@Param			bookId	path	int							true	"code of Book"
//	@Param			request	body	request.UpdateBookRequest	true	"Request of Creating Book Object"
//	@Success		200		{json}	json
//	@Router			/api/book/{bookId} [patch]
func (bookApp *BookApp) Update(ctx *fiber.Ctx) error {

	updateBookRequest := request.UpdateBookRequest{}
	err := ctx.BodyParser(&updateBookRequest)
	handlers.ThrowError(err)

	bookId := ctx.Params("bookId")
	id, err := strconv.Atoi(bookId)
	handlers.ThrowError(err)

	updateBookRequest.Id = id
	bookApp.bookService.Update(updateBookRequest)

	webResponse := response.Response{
		Message: "Successfully updated a book",
	}

	go cacheDelete()


	return ctx.Status(fiber.StatusOK).JSON(webResponse)

}
// DeleteBook Delete Book by Code
//	@Summary		Delete Book by Code
//	@Description	Delete Book by Code in detail
//	@Tags			Books
//	@Accept			json
//	@Produce		json
//	@Param			bookId	path	int	true	"code of Book"
//	@Success		200		{json}	json
//	@Router			/api/book/{bookId} [delete]
func (bookApp *BookApp) Delete(ctx *fiber.Ctx) error {

	bookId := ctx.Params("bookId")
	id, err := strconv.Atoi(bookId)
	handlers.ThrowError(err)
	bookApp.bookService.Delete(id)

	webResponse := response.Response{
		Message: "Successfully deleted a book",
	}

	go cacheDelete()
	

	return ctx.Status(fiber.StatusOK).JSON(webResponse)

}

// FindBook Find Book by Code
//	@Summary		Find Book by Code
//	@Description	Find Book by Code in detail
//	@Tags			Books
//	@Accept			json
//	@Produce		json
//	@Param			bookId	path	int	true	"code of Book"
//	@Success		200		{json}	json
//	@Router			/api/book/{bookId} [get]
func (bookApp *BookApp) FindById(ctx *fiber.Ctx) error {

	bookId := ctx.Params("bookId")
	id, err := strconv.Atoi(bookId)
	handlers.ThrowError(err)
	bookResponse := bookApp.bookService.FindById(id)
	webResponse := response.Response{
	}

	if bookResponse.Name != "" {
		webResponse.Data=bookResponse
	}else{
		result := fmt.Sprintf("Book with id %d not found",id)
		webResponse.Message=result
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)

}

// FindAllBook Find All Book by Code
//	@Summary		Find All Book by Code
//	@Description	Find Book All by Code in detail
//	@Tags			Books
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	json
//	@Router			/api/book [get]
func (bookApp *BookApp) FindAll(ctx *fiber.Ctx) error {

	var bookResponse []response.BooksResponse
	booksCache := cacheRedis.Get("books")
	webResponse := response.Response{
	}

	if booksCache !=nil {
		err := json.Unmarshal(booksCache,&bookResponse)

		if err != nil {
			return nil
		}
		webResponse.Data=bookResponse
	}else{

		bookResponse = bookApp.bookService.FindAll()

		booksListBytes , _ := json.Marshal(bookResponse)

		if len(bookResponse) > 0 {

				go func(books []byte) {
					cacheRedis.Set("books",books)
				}(booksListBytes)

				webResponse.Data= bookResponse
		}else{
			webResponse.Message="Not found books"
		}
	}


	return ctx.Status(fiber.StatusOK).JSON(webResponse)

}
