package main

import (
	"library/controller"
	"library/database"
	"library/repository"
	"library/routes"
	"library/service"
	"log"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/gofiber/swagger"
	_"library/docs"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

//	@title			Library API
//	@version		1.0
//	@description	This is a sample swagger for Library API
//	@termsOfService	http://swagger.io/terms/
//	@contact.name	API Support
//	@contact.email	fiber@swagger.io
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			localhost:3000
//	@BasePath		/
func main() {


	if err := godotenv.Load(); err!= nil {
		log.Println("No .env ile found")
	}

	connect ,err:= database.ConnectDB()

	if err != nil {
		panic("Database is not connection")
	}

	defer connect.DBClose()

	validate := validator.New()

	connect.Runmigrate("books")

	bookRepository := repository.NewBookRepositoryImpl(connect.Db)

	bookService := service.NewBookServiceImpl(bookRepository, validate)

	bookApp := controller.NewBookApp(bookService)

	routes := routes.NewRouter(bookApp)
	
	app := fiber.New()

	app.Use(recover.New())

	app.Get("/swagger/*",swagger.HandlerDefault)

	app.Mount("/api", routes)

	log.Fatal(app.Listen(":3000"))

}
