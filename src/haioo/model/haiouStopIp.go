package model

type HaiouStopIp struct {
	Id          int    `xorm:"not null pk autoincr INT(11)"`
	Ip          string `xorm:"not null VARCHAR(25)"`
	Reason      string `xorm:"default '' VARCHAR(50)"`
	Optime      int    `xorm:"INT(12)"`
	Stoptime    int    `xorm:"INT(12)"`
	Autorelease int    `xorm:"INT(1)"`
	Status      int    `xorm:"not null default 1 TINYINT(1)"`
	Type        int    `xorm:"default 1 TINYINT(1)"`
}
