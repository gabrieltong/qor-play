package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Game struct {
	gorm.Model

	PlayId int
	Play   Play

	UserId int
	User   User

	TurnId    int
	TurnIndex int

	Type               string
	ActorType          string
	StartType          string
	IsPublic           bool
	Code               string
	State              string
	SubState           string
	Round              int
	Size               int
	IsOver             bool
	Conclusion         string `sql:"type:text"`
	KillerWin          bool
	ActivateAutoCancel bool
	Level              int
	Full               int
	lookers            int
	PmType             int
	NoTimeout          int
	TmWeIntro          int

	PlayersCount       int
	TurnTimeoutAt      *time.Time
	TimeoutAt          *time.Time
	TimeoutLimitAt     *time.Time
	StartedAt          *time.Time
	ReadingDoneAt      *time.Time
	IntroductionDoneAt *time.Time
	OverAt             *time.Time
	PiaDoneAt          *time.Time
	FinalDiscussDoneAt *time.Time
	CancelTimeoutAt    *time.Time
	CancelAt           *time.Time
	RequestCancelAt    *time.Time

	Discuss1DoneAt *time.Time
	Discuss2DoneAt *time.Time
	Discuss3DoneAt *time.Time
	Discuss4DoneAt *time.Time

	Round1DoneAt *time.Time
	Round2DoneAt *time.Time
	Round3DoneAt *time.Time
	Round4DoneAt *time.Time

	DiscussAll1DoneAt *time.Time
	DiscussAll2DoneAt *time.Time
	DiscussAll3DoneAt *time.Time
	DiscussAll4DoneAt *time.Time
}
