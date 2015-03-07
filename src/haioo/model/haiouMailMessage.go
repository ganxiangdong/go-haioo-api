package model

import (
	"time"
)

type HaiouMailMessage struct {
	Id          int       `xorm:"not null pk autoincr INT(8)"`
	Touserid    int       `xorm:"INT(8)"`
	Fromuserid  int       `xorm:"INT(8)"`
	Frominfo    string    `xorm:"default '0' VARCHAR(250)"`
	Msgtype     int       `xorm:"default 1 TINYINT(1)"`
	Sub         string    `xorm:"VARCHAR(50)"`
	Con         string    `xorm:"TEXT"`
	Iflook      string    `xorm:"VARCHAR(10)"`
	Date        time.Time `xorm:"not null default '0000-00-00 00:00:00' DATETIME"`
	Contype     int       `xorm:"TINYINT(4)"`
	Tid         string    `xorm:"VARCHAR(50)"`
	ReceiveType string    `xorm:"VARCHAR(200)"`
	ReplyBy     int       `xorm:"INT(11)"`
	Attachments string    `xorm:"VARCHAR(50)"`
	IsSave      int       `xorm:"default 0 TINYINT(1)"`
}
