package models

import (
	"github.com/gabrieltong/qor-play/config"
	"github.com/jinzhu/gorm"
	"github.com/qor/media/oss"
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
	PlayType               config.Options `sq:"type:text"`
	PlayLevel              config.Options `sq:"type:text"`
	PlaySize               config.Options `sq:"type:text"`
	PlaySizeRange          config.Options `sq:"type:text"`
	Kefu                   string
	IosVersion             string
	IosUrl                 string
	AndroidVersion         string
	AndroidUrl             string

	AboutFile     oss.OSS
	QrcodeFile    oss.OSS
	GuideFile     oss.OSS
	AgreementFile oss.OSS
}
