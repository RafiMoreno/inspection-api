package models

type ImageField struct {
	ID uint `json:"id" gorm:"primaryKey;autoIncrement"`
	Label string `json:"label"`
	ImageUrl string `json:"image_url"`
}