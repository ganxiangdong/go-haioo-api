package model

import (
	"time"
)

type HaiouAnnouncement struct {
	Id           int       `xorm:"not null pk autoincr INT(11)"`
	Title        string    `xorm:"not null VARCHAR(100)"`
	Content      string    `xorm:"not null TEXT"`
	Url          string    `xorm:"VARCHAR(100)"`
	Status       int       `xorm:"not null default 0 TINYINT(1)"`
	Displayorder int       `xorm:"not null default 0 SMALLINT(6)"`
	CreateUser   string    `xorm:"VARCHAR(30)"`
	CreateTime   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
