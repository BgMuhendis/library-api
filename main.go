package main

import (
	"library/app"
	"library/config"
	"library/repository"
	"library/router"
	"library/service"
	"log"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)



func main() {
	
	if err := godotenv.Load(); err!= nil {
		log.Println("No .env ile found")
	}

	connect := config.ConnectDB()

	defer connect.DBClose()

	validate := validator.New()

	connect.Automigrate("books")

	bookRepository := repository.NewBookRepositoryImpl(connect.Db)

	bookService := service.NewBookServiceImpl(bookRepository, validate)

	bookApp := app.NewBookApp(bookService)

	routes := router.NewRouter(bookApp)

	app := fiber.New()

	app.Mount("/api", routes)

	log.Fatal(app.Listen(":3000"))

}
