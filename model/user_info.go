package model

type Interest int

const (
	LOLITA Interest = iota
	JK
	HANFU
	ALL
)

type UserInfo struct {
	PhoneNum string

	UserToken string

	UserId int

	NickName string

	City string

	Preference Interest

	LoginTimes int

	IdolNum int

	FansNum int
}
