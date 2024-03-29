package main

import (
	"AudioTranscription/serve/controllers"
	"AudioTranscription/serve/db"
	"AudioTranscription/serve/repository"
	"AudioTranscription/serve/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicln(err)
	}
}

func main() {
	conn := db.NewConnection()
	defer conn.Close()

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Hello World"})
	})

	usersRepo := repository.NewUsersRepository(conn)
	authController := controllers.NewAuthController(usersRepo)
	authRoutes := routes.NewAuthRoutes(authController)
	authRoutes.Install(app)
	// Obtener todas las rutas
	routes := app.GetRoutes()

	// Imprimir todas las rutas
	fmt.Println("Rutas registradas:")
	for _, route := range routes {
		fmt.Printf("%s %s\n", route.Method, route.Path)
	}
	log.Fatal(app.Listen(":8080"))
}
