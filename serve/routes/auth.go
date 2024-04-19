package routes

import (
	"AudioTranscription/serve/controllers"
	"github.com/gofiber/fiber/v2"
)

type authRoutes struct {
	authController controllers.AuthController
}

func NewAuthRoutes(authController controllers.AuthController) Routes {
	return &authRoutes{authController}
}

func (r *authRoutes) Install(app *fiber.App) {

	auth(app, r)
	users(app, r)
}

func auth(app *fiber.App, r *authRoutes) {
	a := app.Group("/auth")
	a.Post("/signup", r.authController.SignUp)
	a.Post("/signin", r.authController.SignIn)

}

func users(app *fiber.App, r *authRoutes) {
	u := app.Group("/users", AuthRequired)
	u.Get("/", r.authController.GetUsers)
	u.Get("/:id", r.authController.GetUser)
	u.Put("/:id", r.authController.PutUser)
	u.Delete("/:id", r.authController.DeleteUser)
}
