package models

import (
	"github.com/jinzhu/gorm"
)

type Actor struct {
	gorm.Model
	Play            Play
	PlayId          int
	IsKiller        bool
	Name            string
	Sex             int
	Age             int
	ShortDesc       string `sql:"type:text"`
	Desc            string `sql:"type:text"`
	DefaultIntro    string
	Point1          int `gorm:"column:point_1"`
	Point2          int `gorm:"column:point_2"`
	Point3          int `gorm:"column:point_3"`
	Point4          int `gorm:"column:point_4"`
	IsNpc           bool
	PublicClueLimit int
	Show            bool

	TaskTpls []TaskTpl
}
