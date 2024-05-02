package models

// JobModel struct
// swagger:model
// @Description Job struct register in the application
type JobModel struct {
	ID         uint   `gorm:"primarykey"`
	Queue      string `gorm:"size:100;not null" json:"queue"`
	Payload    string `gorm:"type:varchar(850);not null" json:"payload"`
	Attempts   uint8  `gorm:"not null" json:"attempts"`
	Completed  bool   `json:"completed"`
	UnResolved bool   `json:"unresolved"`
	CreatedAt  uint   `gorm:"not null" json:"created_at"`
}
