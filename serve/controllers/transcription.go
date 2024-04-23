package controllers

import (
	"AudioTranscription/serve/repository"
	"github.com/gofiber/fiber/v2"
)

type TranscriptionController interface {
	// swagger:route GET /transcription Transcription listTranscription
	// Returns a list of transcriptions
	// responses:
	// 	200: transcriptionResponse
	ListTranscription(ctx *fiber.Ctx) error

	// swagger:route POST /transcription Transcription createTranscription
	// Create a new transcription
	// responses:
	// 	201: transcriptionResponse
	CreateTranscription(ctx *fiber.Ctx) error

	// swagger:route GET /transcription/{id} Transcription getTranscription
	// Get a transcription by id
	// responses:
	// 	200: transcriptionResponse
	GetTranscription(ctx *fiber.Ctx) error

	// swagger:route PUT /transcription/{id} Transcription updateTranscription
	// Update a transcription by id
	// responses:
	// 	200: transcriptionResponse
	UpdateTranscription(ctx *fiber.Ctx) error

	// swagger:route DELETE /transcription/{id} Transcription deleteTranscription
	// Delete a transcription by id
	// responses:
	// 	204: noContentResponse
	DeleteTranscription(ctx *fiber.Ctx) error
}

type transcriptionController struct {
	transcriptionRepo repository.TranscriptionRepository
}

func (t transcriptionController) ListTranscription(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (t transcriptionController) CreateTranscription(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (t transcriptionController) GetTranscription(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (t transcriptionController) UpdateTranscription(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (t transcriptionController) DeleteTranscription(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func NewTranscriptionController(transcription repository.TranscriptionRepository) TranscriptionController {
	return &transcriptionController{transcription}
}
