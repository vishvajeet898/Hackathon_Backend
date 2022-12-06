package model

import "github.com/lib/pq"

type Basic_Info struct {
	BasicId    string         `json:"basic_id" gorm:"primaryKey" gorm:"column:user_id" validate:"required"`
	UserId     string         `json:"user_id" gorm:"column:user_id"`
	Age        int            `json:"age" gorm:"column:age"`
	Height     string         `json:"height" gorm:"column:height"`
	Sex        string         `json:"sex" gorm:"column:sex"`
	BloodGroup string         `json:"blood_group" gorm:"column:blood_group"`
	Weight     int            `json:"weight" gorm:"column:weight"`
	Allergies  pq.StringArray `json:"allergies" gorm:"type:text[]" gorm:"column:allergies"`
	Diseases   pq.StringArray `json:"diseases" gorm:"type:text[]" gorm:"column:diseases"`
}
