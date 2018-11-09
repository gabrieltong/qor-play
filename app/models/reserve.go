package models

import (
	"github.com/gabrieltong/qor-play/config"
)

type Reserve struct {
	config.Model
	UserId int
	User   User
	PlayId int
	Play   Play
}
