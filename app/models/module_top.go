package models

import (
	"github.com/jinzhu/gorm"
)

type ModuleTop struct {
	gorm.Model
	Title     string
	Link      string
	Position  int
	InnerLink bool
}
