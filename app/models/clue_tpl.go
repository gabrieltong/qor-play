package models

import (
	"github.com/jinzhu/gorm"
)

type ClueTpl struct {
	gorm.Model

	ClueMenuId int
	ClueMenu   ClueMenu

	Title     string
	LockTitle string
	Point     int
	Round     int
}
