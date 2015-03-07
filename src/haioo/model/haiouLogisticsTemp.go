package model

type HaiouLogisticsTemp struct {
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	Title      string `xorm:"VARCHAR(255)"`
	Type       int    `xorm:"default 0 TINYINT(1)"`
	Desc       string `xorm:"TEXT"`
	Expresstpl string `xorm:"VARCHAR(20)"`
	Status     int    `xorm:"default 1 TINYINT(1)"`
}
