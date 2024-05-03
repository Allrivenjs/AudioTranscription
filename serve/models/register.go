package models

import "AudioTranscription/serve/db"

func AutoMigrate(conn db.Connection) {
	conn.RegisterModels(
		&User{},
		&Transcription{},
		&JobModel{},
	)
	conn.Migrate()
}
