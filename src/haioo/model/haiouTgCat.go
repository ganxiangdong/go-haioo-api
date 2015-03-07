package model

type HaiouTgCat struct {
	Id           int    `xorm:"not null pk autoincr INT(6)"`
	ParentId     int    `xorm:"not null default 0 INT(11)"`
	Catname      string `xorm:"not null VARCHAR(30)"`
	Displayorder int    `xorm:"not null default 255 SMALLINT(8)"`
}
