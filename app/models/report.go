package models

import (
	"github.com/gabrieltong/qor-play/config"
)

type Report struct {
	config.Model
	UserId  int
	User    User
	FromId  int
	Content string `sql:"type:text"`
}
