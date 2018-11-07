package models

import (
	"github.com/jinzhu/gorm"
)

type Activity struct {
	*gorm.Model
	trackable_type string `gorm:"index:trackable_type"`
	trackable_id   int    `gorm:"index:trackable_id"`
	owner_type     string `gorm:"index:owner_type"`
	owner_id       int    `gorm:"index:owner_id"`
	key            string `gorm:"index:key"`
	parameters     string
	recipient_type string `gorm:"index:recipient_type"`
	recipient_id   int    `gorm:"index:recipient_id"`
}
