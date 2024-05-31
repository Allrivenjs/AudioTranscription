package main

import (
	_ "AudioTranscription/docs"
	"AudioTranscription/serve/controllers"
	"AudioTranscription/serve/db"
	"AudioTranscription/serve/jobs"
	"AudioTranscription/serve/models"
	"AudioTranscription/serve/repository"
	"AudioTranscription/serve/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

const loadEnv = true

func init() {
	if loadEnv {
		err := godotenv.Load()
		if err != nil {
			log.Panicln(err)
		}
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
	//getRoutes := app.GetRoutes()

	//Imprimir todas las rutas
	//fmt.Println("Rutas registradas:")
	//for _, route := range getRoutes {
	//	fmt.Printf("-> %s %s\n", route.Method, route.Path)
	//}
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
	// The key for the map is message.to
	//clients := make(map[string]string)

	app := fiber.New(fiber.Config{
		BodyLimit:         1024 * 1024 * 1024,
		StreamRequestBody: true,
	})

	// public storage
	app.Static("/storage", "./storage")

	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Hello World", "os": runtime.GOOS, "distro": runtime.GOARCH})
	})
	app.Get("/healthz", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "pong"})
	})
	app.Get("/install", func(ctx *fiber.Ctx) error {
		if runtime.GOOS == "linux" {
			cmd := exec.Command("apt-get", "install", "ffmpeg")
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			if err := cmd.Run(); err != nil {
				fmt.Println("Error al instalar ffmpeg:", err)
				return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Error al instalar ffmpeg", "Error": err.Error()})
			}
		}
		return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "ffmpeg instalado"})
	})

	r := &appRepository{app: app}
	conn := r.async()

	app.Get("/reset", func(ctx *fiber.Ctx) error {
		conn.RefreshMigration()
		return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Base de datos reiniciada"})
	})

	defer conn.Close()
	go jobs.Init(conn).Listen()
	fmt.Println(os.Getenv("APP_PORT"))
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))

}
