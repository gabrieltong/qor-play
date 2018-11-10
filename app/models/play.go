package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/qor/media"
	"github.com/qor/media/oss"
)

var PlayState []string = []string{"draft", "approving", "approved", "rejected", "published", "noticed"}
var ClueTplType = [][]string{{"1", "菜单"}, {"2", "图片"}}
var PlayType = [][]string{{"1", "恐怖"}, {"2", "惊悚"}, {"3", "热点"}, {"4", "本格"}, {"5", "童话"}, {"6", "宫斗"}, {"7", "搞笑"}, {"8", "武侠"}, {"9", "现代"}, {"10", "民国"}, {"11", "怀旧"}}
var PlayLevel = [][]string{{"1", "萌新"}, {"2", "普通"}, {"3", "烧脑"}}
var PlayRoundSize = [][]string{{"2", "2轮"}, {"3", "3轮"}, {"4", "4轮"}}
var PlaySize = [][]string{{"1", "1"}, {"2", "2"}, {"3", "3"}, {"4", "4"}, {"5", "5"}, {"6", "6"}, {"7", "7"}}

type Play struct {
	gorm.Model

	Actors            []Actor
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
	RongNo            string

	TmIntro         int
	TmRead          int
	TmRound         int
	TmPoll          int
	TmTotal         int
	TmDiscuss       int
	TmIntroPlayer   int
	TmIntroAll      int
	TmDiscussPlayer int
	TmDiscussAll    int
	TmOver          int
	TmPia           int

	TmOverLimit          int
	TmIntroPlayerLimit   int
	TmIntroAllLimit      int
	TmDiscussPlayerLimit int
	TmDiscussAllLimit    int
	TmTotalLimit         int
	TmReadLimit          int
	TmIntroLimit         int
	TmRoundLimit         int
	TmDiscussLimit       int
	TmPollLimit          int
	TmPiaLimit           int

	GamesWaitingCount      int
	GamesReadingCount      int
	GamesPiaCount          int
	GamesIntroductionCount int

	GamesDiscuss1Count int `gorm:"column:games_discuss_1_count"`
	GamesDiscuss2Count int `gorm:"column:games_discuss_2_count"`
	GamesDiscuss3Count int `gorm:"column:games_discuss_3_count"`
	GamesDiscuss4Count int `gorm:"column:games_discuss_4_count"`

	GamesDiscussAll1Count int `gorm:"column:games_discuss_all_1_count"`
	GamesDiscussAll2Count int `gorm:"column:games_discuss_all_2_count"`
	GamesDiscussAll3Count int `gorm:"column:games_discuss_all_3_count"`
	GamesDiscussAll4Count int `gorm:"column:games_discuss_all_4_count"`

	GamesRound1Count int `gorm:"column:games_round_1_count"`
	GamesRound2Count int `gorm:"column:games_round_2_count"`
	GamesRound3Count int `gorm:"column:games_round_3_count"`
	GamesRound4Count int `gorm:"column:games_round_4_count"`

	GamesFinalDiscussCount int
	GamesPollingCount      int
	GamesOverCount         int
	GamesCanceledCount     int

	Debug     bool
	NoTimeout bool

	RequestCancelTimeout int
	CancelNeed           int

	ReservesCount int
	PlayersCount  int
	Conclusion    string `sql:"type:text"`
	DescPre1      string `gorm:"column:desc_pre_1" sql:"type:text"`
	DescPre2      string `gorm:"column:desc_pre_2" sql:"type:text"`

	ClueTplType int

	// Opening oss.OSS
}

type Opening struct{ oss.OSS }

func (Opening) GetSizes() map[string]*media.Size {
	return map[string]*media.Size{
		"small":  {Width: 50, Height: 50},
		"middle": {Width: 120, Height: 120},
		"big":    {Width: 320, Height: 320},
	}
}
