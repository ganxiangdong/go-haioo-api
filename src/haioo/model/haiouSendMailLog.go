package model

import (
	"time"
)

type HaiouSendMailLog struct {
	Id         int       `xorm:"not null pk autoincr INT(10)"`
	Mail       string    `xorm:"not null default '' index CHAR(32)"`
	Subject    string    `xorm:"not null default '' VARCHAR(100)"`
	Content    string    `xorm:"TEXT"`
	CreateTime time.Time `xorm:"not null index DATETIME"`
	Status     int       `xorm:"not null default 1 index TINYINT(1)"`
}
