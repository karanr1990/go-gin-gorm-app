package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Tittle string
	Body   string
}
