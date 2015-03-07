package model

import (
	"time"
)

type HaiouSns struct {
	Id                   int       `xorm:"not null pk autoincr INT(11)"`
	MemberId             int       `xorm:"not null INT(11)"`
	Title                string    `xorm:"VARCHAR(500)"`
	Content              string    `xorm:"not null TEXT"`
	Pic                  string    `xorm:"VARCHAR(500)"`
	Status               int       `xorm:"not null default 1 TINYINT(1)"`
	Privacy              int       `xorm:"not null default 0 TINYINT(1)"`
	Tags                 string    `xorm:"default '' VARCHAR(200)"`
	CommentCount         int       `xorm:"not null default 0 INT(11)"`
	CopyCount            int       `xorm:"not null default 0 INT(11)"`
	DigsCount            int       `xorm:"not null default 0 INT(11)"`
	OriginalId           int       `xorm:"not null default 0 INT(11)"`
	OriginalMemberId     int       `xorm:"not null default 0 INT(11)"`
	OriginalStatus       int       `xorm:"not null default 1 TINYINT(1)"`
	OriginalCommentCount int       `xorm:"not null default 0 INT(11)"`
	OriginalDigsCount    int       `xorm:"not null default 0 INT(11)"`
	OriginalCopyCount    int       `xorm:"not null default 0 INT(11)"`
	ClientPlatform       int       `xorm:"TINYINT(1)"`
	ClientSystem         int       `xorm:"TINYINT(1)"`
	CreateTime           time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
