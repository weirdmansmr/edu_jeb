package models

import (
	"github.com/jinzhu/gorm"
)

type Lesson struct {
    gorm.Model
    Title   string `json:"title"`
    Content string `json:"content"`
}