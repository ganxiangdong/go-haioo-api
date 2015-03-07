package model

type HaiouMemberCount struct {
	MemberId int `xorm:"not null pk INT(11)"`
	Growth   int `xorm:"default 0 INT(11)"`
	Points   int `xorm:"default 0 INT(11)"`
}
