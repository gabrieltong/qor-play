package models

import (
	"github.com/jinzhu/gorm"
)

type TaskTpl struct {
	gorm.Model
	Title string
	Score int
	Type  int

	ActorId int
	Actor   Actor

	TaskTplAnswers []TaskTplAnswer
}
