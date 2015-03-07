package model

import (
	"time"
)

type HaiouProperty struct {
	Id           int       `xorm:"not null pk autoincr INT(11)"`
	Name         string    `xorm:"VARCHAR(100)"`
	Displayorder int       `xorm:"not null default 0 SMALLINT(6)"`
	CreateUser   string    `xorm:"VARCHAR(30)"`
	CreateTime   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
