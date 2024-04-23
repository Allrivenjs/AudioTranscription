package models

import "gorm.io/gorm"

// Transcription struct
// swagger:model
// @Description Transcription struct register in the application
type Transcription struct {
	gorm.Model
	UserId            uint   `gorm:"not null" json:"user_id"`
	AudioUrl          string `gorm:"size:255;not null" json:"audio_url"`
	Transcription     string `gorm:"type:text;not null" json:"transcription"`
	SortTranscription string `gorm:"size:255;not null" json:"sort_transcription"`
}
