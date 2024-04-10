package router

import (
	"github.com/gofiber/fiber/v2"
	"library/app"
)

func NewRouter(bookApp *app.BookApp) *fiber.App {
	router := fiber.New()

	router.Get("/check", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"Status":  "Success",
			"Message": "Welcome new project",
		})
	})
	router.Route("/book", func(router fiber.Router) {
		router.Post("/", bookApp.Create)
		router.Get("/", bookApp.FindAll)

	})

	router.Route("/book/:bookId", func(router fiber.Router) {
		router.Delete("/", bookApp.Delete)
		router.Get("/", bookApp.FindById)
		router.Patch("/", bookApp.Update)

	})

	return router
}
