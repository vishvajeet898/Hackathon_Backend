package model

import "time"

type Record struct {
	RecordId  string    `json:"file_id" gorm:"primaryKey" gorm:"column:user_id" validate:"required"`
	UserId    string    `json:"user_id" gorm:"column:user_id"`
	AwsLink   string    `json:"aws_link" gorm:"column:type"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}
