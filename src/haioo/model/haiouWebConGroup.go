package model

type HaiouWebConGroup struct {
	Id    int    `xorm:"not null pk autoincr INT(11)"`
	Title string `xorm:"VARCHAR(60)"`
	Lang  string `xorm:"VARCHAR(5)"`
	Sort  int    `xorm:"default 0 SMALLINT(4)"`
	Logo  string `xorm:"VARCHAR(100)"`
}
