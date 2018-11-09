package models

import (
	"github.com/jinzhu/gorm"
)

type Setting struct {
	gorm.Model
	Title                  string
	SExp                   int
	AExp                   int
	BExp                   int
	CExp                   int
	LevelExp               int
	ExpLast                int
	CheckExp               int
	HighLevel              int
	AutoCancelOnwerOffline int
	PlayType               string
	PlayLevel              string
	PlaySize               string
	PlaySizeRange          string
	Kefu                   string
	IosVersion             string
	IosUrl                 string
	AndroidVersion         string
	AndroidUrl             string
}
