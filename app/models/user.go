package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name                 string
	Level                int
	PlayReservedCount    int
	PlayPlayedCount      int
	PlayOwned            int
	GameOwnedCount       int
	Agent                string
	Unionid              string
	Openid               string
	Headimgurl           string
	Raw                  string `sql:"type:text"`
	Rtoken               string
	Password             string
	PasswordConfirmation string
	PasswordDigest       string
	PlayersCount         int
	ReservesCount        int
	RongId               string
	Exp                  int
	ExpLast              int
	Sex                  int
	Age                  int
	City                 string
	PlayedCount          int
	LettersCount         int
	UnreadLettersCount   int
	ReadedLettersCount   int
	Online               bool
	ShowId               int
}
