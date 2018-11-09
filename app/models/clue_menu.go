package models

import (
	"github.com/jinzhu/gorm"
)

type ClueMenu struct {
	gorm.Model

	PlayId int
	Play   Play

	ClueMenus []ClueMenu
	// LockById int
	// LockBy   ClueTpl

	Title string
	Point int
	Round int
	Size  int
}
