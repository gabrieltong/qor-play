package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Play struct {
	gorm.Model
	Title             string
	Writer            string
	OnlineAt          time.Time
	Desc              string `sql:"type:text"`
	ActorDesc         string `sql:"type:text"`
	ShortDesc         string `sql:"type:text"`
	AdminUserId       int
	AdminUser         AdminUser
	Duration          int
	Type              int
	State             string
	Size              int
	Score             int
	RoundSize         int
	RoundPoint        int
	Level             int
	GamesCount        int
	UserReservedCount int

	TmIntro       int
	TmRead        int
	TmRound       int
	TmPoll        int
	TmTotal       int
	TmDiscuss     int
	TmIntroPlay   int
	TmIntroAll    int
	TmDiscussPlay int
	TmDiscussAll  int
	TmOver        int
	TmPia         int

	TmOverLimit          int
	TmIntroPlayerLimit   int
	TmIntroAllLimit      int
	TmDiscussPlayerLimit int
	TmDisscussAllLimit   int
	TmTotalLimit         int
	TmReadLimit          int
	TmIntroLimit         int
	TmRoundLimit         int
	TmDiscussLimit       int
	TmPollLimit          int
	TmPiaLimit           int

	GamesWaitingCount      int
	GamesReadingCount      int
	GamesPicCount          int
	GamesIntroductionCount int

	GamesDiscuss1Count int
	GamesDiscuss2Count int
	GamesDiscuss3Count int
	GamesDiscuss4Count int

	GamesDiscussAll1Count int
	GamesDiscussAll2Count int
	GamesDiscussAll3Count int
	GamesDiscussAll4Count int

	GamesRound1Count int
	GamesRound2Count int
	GamesRound3Count int
	GamesRound4Count int

	GamesFinalDiscussCount int
	GamesPollingCount      int
	GamesOverCount         int
	GamesCanceledCount     int

	Debug     bool
	NoTimeout bool

	RongNo               string
	RequestCancelTimeout int
	CancelNeed           int

	ReservesCount int
	PlayersCount  string
	Conclusion    string `sql:"type:text"`
	DescPre1      string `sql:"type:text"`
	DescPre2      string `sql:"type:text"`

	CluTplType int
}
