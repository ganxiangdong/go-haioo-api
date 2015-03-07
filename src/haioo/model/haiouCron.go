package model

type HaiouCron struct {
	Id           int    `xorm:"not null pk autoincr INT(6)"`
	Name         string `xorm:"VARCHAR(50)"`
	Script       string `xorm:"VARCHAR(50)"`
	Lasttransact int    `xorm:"INT(10)"`
	Nexttransact int    `xorm:"INT(10)"`
	Week         string `xorm:"default '-1' VARCHAR(12)"`
	Day          string `xorm:"default '-1' VARCHAR(2)"`
	Hours        string `xorm:"default '00' VARCHAR(2)"`
	Minutes      string `xorm:"default '00' VARCHAR(2)"`
	Everyminute  int    `xorm:"INT(10)"`
	Active       int    `xorm:"default 0 TINYINT(1)"`
}
