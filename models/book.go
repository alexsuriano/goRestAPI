package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Author      string         `json:"author"`
	Description string         `json:"description"`
	MediumPrice float32        `json:"medium_price"`
	ImageURL    string         `json:"img_url"`
	CreatedAt   time.Time      `json:"created"`
	UpdateAT    time.Time      `json:"updated"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deteted"`
}
