package models

import (
	"github.com/jinzhu/gorm"
)

type Title struct {
	gorm.Model
	Title    string
	Type     string
	Desc     string `sql:"type:text"`
	Countent string `sql:"type:text"`
	Level    int
}
