package model

import (
	"time"
)

type HaiouSendMsgLog struct {
	Id         int       `xorm:"not null pk autoincr INT(10)"`
	Mobile     string    `xorm:"not null default '' index CHAR(11)"`
	Content    string    `xorm:"not null default '' VARCHAR(80)"`
	CreateTime time.Time `xorm:"not null DATETIME"`
	Status     int       `xorm:"not null default 1 TINYINT(1)"`
}
