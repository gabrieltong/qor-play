package models

import (
	"github.com/jinzhu/gorm"
)

type Player struct {
	gorm.Model

	ActorId int
	Actor   Actor

	UserId int
	User   User

	GameId int
	Game   Game

	PlayId int
	Play   Play

	KillerActorIds string

	ClueSize int
	Ready    bool

	ReadingDone      bool
	IntroductionDone bool
	Round1Done       bool
	Round2Done       bool
	Round3Done       bool
	Round4Done       bool
	Discuss1Done     bool
	Discuss2Done     bool
	Discuss3Done     bool
	Discuss4Done     bool
	PollDone         bool

	Introduction string `sql:"type:text"`
	Win          bool
}
