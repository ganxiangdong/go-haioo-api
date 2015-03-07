package model

import (
	"time"
)

type HaiouMemberOpinions struct {
	Id             int       `xorm:"not null pk autoincr INT(8)"`
	MemberId       int       `xorm:"not null index INT(8)"`
	OpinionType    int       `xorm:"not null default 0 index TINYINT(1)"`
	OpinionContent string    `xorm:"TEXT"`
	Contact        string    `xorm:"default '' VARCHAR(30)"`
	CreateTime     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	ReplyTime      time.Time `xorm:"TIMESTAMP"`
	Status         int       `xorm:"default 0 TINYINT(1)"`
}
