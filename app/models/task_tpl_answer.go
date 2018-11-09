package models

import (
	"github.com/jinzhu/gorm"
)

type TaskTplAnswer struct {
	gorm.Model
	TaskTplId int
	TaskTpl   TaskTpl
	Title     string
	IsAnswer  bool
}
