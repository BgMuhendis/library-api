package main

import (
	"library/app"
	"library/config"
	"library/model"
	"library/repository"
	"library/router"
	"library/service"
	"log"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)



func main() {
	loadConfig, err := config.EnvLoad(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	db := config.ConnectDB(&loadConfig)
	validate := validator.New()

	db.Table("books").AutoMigrate(&model.Book{})

	bookRepository := repository.NewBookRepositoryImpl(db)

	bookService := service.NewBookServiceImpl(bookRepository, validate)

	bookApp := app.NewBookApp(bookService)

	routes := router.NewRouter(bookApp)

	app := fiber.New()

	app.Mount("/api", routes)

	log.Fatal(app.Listen(":3000"))

}
