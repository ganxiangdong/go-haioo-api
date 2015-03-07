package model

import (
	"time"
)

type HaiouVote struct {
	Id       int       `xorm:"not null pk autoincr INT(11)"`
	Newsid   int       `xorm:"not null default 0 SMALLINT(6)"`
	Title    string    `xorm:"not null VARCHAR(120)"`
	Votetext string    `xorm:"not null TEXT"`
	Votetype int       `xorm:"not null default 0 TINYINT(1)"`
	Num      int       `xorm:"not null INT(6)"`
	Limitip  int       `xorm:"not null default 0 TINYINT(1)"`
	Time     time.Time `xorm:"not null default '0000-00-00' DATE"`
	Tempid   int       `xorm:"not null default 0 SMALLINT(6)"`
	Type     int       `xorm:"TINYINT(4)"`
	Uptime   int       `xorm:"not null INT(10)"`
}
