package repository

import (
	"AudioTranscription/serve/db"
	"AudioTranscription/serve/models"
	"gorm.io/gorm"
)

type TranscriptionRepository interface {
	SaveOrUpdate(transcription *models.Transcription) error
	GetById(id string) (*models.Transcription, error)
	GetAll() ([]*models.Transcription, error)
	Delete(id string) error
}

type transcriptionRepository struct {
	db *gorm.DB
}

func (t transcriptionRepository) SaveOrUpdate(transcription *models.Transcription) error {
	//TODO implement me
	panic("implement me")
}

func (t transcriptionRepository) GetById(id string) (*models.Transcription, error) {
	//TODO implement me
	panic("implement me")
}

func (t transcriptionRepository) GetAll() ([]*models.Transcription, error) {
	//TODO implement me
	panic("implement me")
}

func (t transcriptionRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewTranscriptionRepository(conn db.Connection) TranscriptionRepository {
	return &transcriptionRepository{db: conn.DB()}
}
