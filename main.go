package main

import (
	_ "AudioTranscription/docs"
	"AudioTranscription/serve/controllers"
	"AudioTranscription/serve/db"
	"AudioTranscription/serve/models"
	"AudioTranscription/serve/repository"
	"AudioTranscription/serve/routes"
	"AudioTranscription/serve/services"
	"AudioTranscription/serve/storage"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
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

type app interface {
	async(app *fiber.App) db.Connection
	registerDocSwagger(app *fiber.App)
}

type appRepository struct {
	app *fiber.App
}

func (a *appRepository) async() db.Connection {
	app := a.app
	conn := db.NewConnection()
	models.AutoMigrate(conn)
	usersRepo := repository.NewUsersRepository(conn)
	authController := controllers.NewAuthController(usersRepo)
	authRoutes := routes.NewAuthRoutes(authController)

	transRepo := repository.NewTranscriptionRepository(conn)
	transController := controllers.NewTranscriptionController(transRepo)
	transRoutes := routes.NewTransRoutes(transController)

	authRoutes.Install(app)
	transRoutes.Install(app)
	a.registerDocSwagger()

	// Obtener todas las rutas
	getRoutes := app.GetRoutes()

	//Imprimir todas las rutas
	fmt.Println("Rutas registradas:")
	for _, route := range getRoutes {
		fmt.Printf("-> %s %s\n", route.Method, route.Path)
	}
	return conn
}

func (a *appRepository) registerDocSwagger() {
	app := a.app
	app.Get("/docs/*", swagger.HandlerDefault)

	app.Get("/docs/*", swagger.New(swagger.Config{
		URL:          "http://example.com/swagger.json",
		DeepLinking:  false,
		DocExpansion: "none",
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		OAuth2RedirectUrl: "http://localhost:3001/swagger/oauth2-redirect.html",
	}))
}

func main() {

	//cloudflare.CloudflareAI()

	app := fiber.New(fiber.Config{
		BodyLimit:         1024 * 1024 * 1024,
		StreamRequestBody: true,
	})

	// public storage
	app.Static("/storage", "./storage")

	app.Use(cors.New())
	app.Use(logger.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		//get param file
		fileP := ctx.Query("file")
		// verify if existed the folder
		err := storage.CreateFolderTemp()
		if err != nil {
			fmt.Printf("Error creating folder: %s", err.Error())
		}
		path := fmt.Sprintf("%s%s%s", storage.GetPathCurrent(), storage.GetBaseRoute(), storage.GetBaseTemp())
		fmt.Println(path)
		_, err = services.CuterAudio(fileP, path)
		if err != nil {
			return err
		}

		return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Hello World", "file": fileP})
	})

	r := &appRepository{app: app}
	conn := r.async()
	defer conn.Close()
	log.Fatal(app.Listen(":8080"))

}
