package routes

import (
	"AudioTranscription/serve/controllers"
	"github.com/gofiber/fiber/v2"
)

type transRoutes struct {
	transController controllers.TranscriptionController
}

func (t *transRoutes) Install(app *fiber.App) {
	transcription(app, t)
}

func NewTransRoutes(transController controllers.TranscriptionController) Routes {
	return &transRoutes{transController}
}

func transcription(app *fiber.App, r *transRoutes) {
	a := app.Group("/transcription")
	a.Get("/", r.transController.ListTranscription)
	a.Post("/", r.transController.CreateTranscription)
	a.Get("/:id", r.transController.GetTranscription)
	a.Put("/:id", r.transController.UpdateTranscription)
	a.Delete("/:id", r.transController.DeleteTranscription)
}
