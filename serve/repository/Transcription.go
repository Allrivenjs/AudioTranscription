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
	if err := t.db.Save(transcription).Error; err != nil {
		return err
	}
	return nil
}

func (t transcriptionRepository) GetById(id string) (*models.Transcription, error) {
	var transcription models.Transcription
	if err := t.db.Find(&transcription, id).Error; err != nil {
		return nil, err
	}
	return &transcription, nil
}

func (t transcriptionRepository) GetAll() ([]*models.Transcription, error) {
	var transcriptions []*models.Transcription
	if err := t.db.Find(&transcriptions).Error; err != nil {
		return nil, err
	}
	return transcriptions, nil
}

func (t transcriptionRepository) Delete(id string) error {
	if err := t.db.Delete(&models.Transcription{}, id).Error; err != nil {
		return err
	}
	return nil
}

func NewTranscriptionRepository(conn db.Connection) TranscriptionRepository {
	return &transcriptionRepository{db: conn.DB()}
}
