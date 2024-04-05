package main

import (
	"AudioTranscription/serve/db"
	"AudioTranscription/serve/models"
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
func async(app *fiber.App) {
	conn := db.NewConnection()
	defer conn.Close()
	models.AutoMigrate(conn)
	//usersRepo := repository.NewUsersRepository(conn)
	//authController := controllers.NewAuthController(usersRepo)
	//authRoutes := routes.NewAuthRoutes(authController)
	//authRoutes.Install(app)

	//// Obtener todas las rutas
	//routes := app.GetRoutes()
	//// get all users
	//app.Get("/userss", authController.GetUsers)
	// Imprimir todas las rutas
	//fmt.Println("Rutas registradas:")
	//for _, route := range routes {
	//	fmt.Printf("-> %s %s\n", route.Method, route.Path)
	//}
}

func main() {

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Hello World"})
	})

	async(app)

	log.Fatal(app.Listen(":8080"))
}
