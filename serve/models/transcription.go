package models

import "gorm.io/gorm"

// Transcription struct
// swagger:model
// @Description Transcription struct register in the application
type Transcription struct {
	gorm.Model
	Title             string `gorm:"size:255;not null" json:"title"`
	UserId            uint   `gorm:"not null" json:"user_id"`
	AudioUrl          string `gorm:"size:255" json:"audio_url"`
	LocateFile        string `gorm:"size:255" json:"locate_file"`
	Transcription     string `gorm:"type:text" json:"transcription"`
	SortTranscription string `gorm:"size:255" json:"sort_transcription"`
}
