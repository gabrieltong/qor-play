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
	Point1          int
	Point2          int
	Point3          int
	Point4          int
	IsNpc           bool
	PublicClueLimit int
	Show            bool
}
