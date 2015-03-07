package model

import (
	"time"
)

type HaiouPointsLog struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	MemberId   int       `xorm:"INT(11)"`
	MemberName string    `xorm:"VARCHAR(100)"`
	Points     int       `xorm:"default 0 INT(11)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Desc       string    `xorm:"VARCHAR(255)"`
}
