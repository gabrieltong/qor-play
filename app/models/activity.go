package models

import (
	"github.com/jinzhu/gorm"
)

type Activity struct {
	gorm.Model
	RrackableType  string `gorm:"index:trackable_type"`
	TrackableId    int    `gorm:"index:trackable_id"`
	OwnerType      string `gorm:"index:owner_type"`
	OwnerId        int    `gorm:"index:owner_id"`
	Key            string `gorm:"index:key"`
	Parameters     string
	Recipient_type string `gorm:"index:recipient_type"`
	Recipient_id   int    `gorm:"index:recipient_id"`
}
