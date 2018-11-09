package models

import (
	"time"

	"github.com/gabrieltong/qor-play/config"
)

type Timeline struct {
	config.Model

	PlayerId int
	Player   Player

	time *time.Time
	desc string `sql:"type:text"`
}
