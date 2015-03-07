package model

import (
	"time"
)

type HaiouFilterKeyword struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	Keyword    string    `xorm:"VARCHAR(50)"`
	Replace    string    `xorm:"VARCHAR(50)"`
	Status     int       `xorm:"default 1 index TINYINT(1)"`
	Time       int       `xorm:"INT(11)"`
	CreateUser string    `xorm:"VARCHAR(30)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
