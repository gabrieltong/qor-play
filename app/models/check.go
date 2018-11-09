package models

import (
	"github.com/jinzhu/gorm"
)

type Check struct {
	gorm.Model

	UserId int
	User   User

	exp int
}
