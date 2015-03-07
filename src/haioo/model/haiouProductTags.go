package model

import (
	"time"
)

type HaiouProductTags struct {
	Id           int       `xorm:"not null pk autoincr INT(9)"`
	Tags         string    `xorm:"VARCHAR(50)"`
	Title        string    `xorm:"TEXT"`
	Keyword      string    `xorm:"TEXT"`
	Description  string    `xorm:"TEXT"`
	Logo         string    `xorm:"VARCHAR(150)"`
	Displayorder int       `xorm:"not null SMALLINT(6)"`
	Status       int       `xorm:"not null TINYINT(1)"`
	CreateUser   string    `xorm:"VARCHAR(30)"`
	CreateTime   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
