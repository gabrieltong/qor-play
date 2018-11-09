package models

type UserLetter struct {
	UserId int
	User   User

	LetterId int
	Letter   Letter

	Readed bool
}
