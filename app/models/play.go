package models

import (
	"time"

	"github.com/jinzhu/gorm"
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
	RongId            string

	TmIntro       int
	TmRead        int
	TmRound       int
	TmPoll        int
	TmTotal       int
	TmDiscuss     int
	TmIntroPlayer int
	TmIntroAll    int
	TmDiscussPlay int
	TmDiscussAll  int
	TmOver        int
	TmPia         int

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

	ClueTplType int
}
