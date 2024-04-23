package controllers

import (
	"AudioTranscription/serve/models"
	"AudioTranscription/serve/repository"
	"AudioTranscription/serve/storage"
	"AudioTranscription/serve/util"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
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
	trans, err := t.transcriptionRepo.GetAll()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(util.SuccessResponse(&fiber.Map{
		"transcriptions": trans,
	}))
}

type CreateTranscriptionRequest struct {
	// In: body
	Audio string `valid:"required" json:"audio"`
	// In: body
	Title string `valid:"required" json:"name"`
}

func (t transcriptionController) CreateTranscription(ctx *fiber.Ctx) error {
	var newTranscription CreateTranscriptionRequest
	err := ctx.BodyParser(&newTranscription)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(util.NewJError(err))
	}
	validateErrors := util.ValidateInput(ctx, newTranscription)
	if validateErrors != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.ErrorResponse(validateErrors))
	}
	var transcriptrion = models.Transcription{
		Title:             newTranscription.Title,
		Transcription:     "",
		SortTranscription: "",
	}
	audio := util.AudioToWav()
	file, err := storage.NewStorage().CreateAudioFile(fmt.Sprintf("%s.%s", newTranscription.Title, "wav"), []byte(newTranscription.Audio))
	if err != nil {
		return err
	}
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
