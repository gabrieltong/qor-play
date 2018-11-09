package models

import (
	"github.com/jinzhu/gorm"
)

type AdminUser struct {
	gorm.Model
	Email             string
	Name              string
	IsAuthor          bool
	IsSuper           bool
	EncryptedPassword string
	Password          string
}
