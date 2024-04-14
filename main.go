package main

import (
	"library/app"
	"library/config"
	"library/models"
	"library/repository"
	"library/router"
	"library/service"
	"log"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)



func main() {
/* 	_, err := config.EnvLoad(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	} */

	if err := godotenv.Load(); err!= nil {
		log.Println("No .env ile found")
	}

	db := config.ConnectDB()

	postgres, _ := db.DB()

	defer postgres.Close()

	validate := validator.New()

	db.Table("books").AutoMigrate(&models.Book{})

	bookRepository := repository.NewBookRepositoryImpl(db)

	bookService := service.NewBookServiceImpl(bookRepository, validate)

	bookApp := app.NewBookApp(bookService)

	routes := router.NewRouter(bookApp)

	app := fiber.New()

	app.Mount("/api", routes)

	log.Fatal(app.Listen(":3000"))

}
