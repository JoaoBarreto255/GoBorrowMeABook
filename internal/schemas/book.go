package schemas

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string
	Isbn        string
	Cover       string
	PublishedAt string
}
