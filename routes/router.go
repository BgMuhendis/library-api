package routes

import (
	"library/app"

	"github.com/gofiber/fiber/v2"

	
)

func NewRouter(bookApp *app.BookApp) *fiber.App {
	router := fiber.New()

	router.Route("/book", func(router fiber.Router) {
		router.Post("/", bookApp.Create)
		router.Get("/", bookApp.FindAll)

	})

	router.Route("/book/:bookId", func(router fiber.Router) {
		router.Delete("/", bookApp.Delete)
		router.Get("/", bookApp.FindById)
		router.Patch("/", bookApp.Update)

	})

	router.Use(func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message":"Endpoint not found!",
		})
	})

	return router
}
