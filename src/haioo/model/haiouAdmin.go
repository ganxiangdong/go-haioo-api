package model

import (
	"time"
)

type HaiouAdmin struct {
	Id           int       `xorm:"not null pk autoincr unique INT(3)"`
	User         string    `xorm:"not null CHAR(30)"`
	Name         string    `xorm:"VARCHAR(50)"`
	Password     string    `xorm:"not null CHAR(35)"`
	GroupId      int       `xorm:"not null default 0 SMALLINT(5)"`
	Desc         string    `xorm:"TEXT"`
	Logonums     int       `xorm:"default 0 INT(5)"`
	Lastlogotime int       `xorm:"INT(11)"`
	Logoip       string    `xorm:"VARCHAR(30)"`
	Province     string    `xorm:"VARCHAR(60)"`
	City         string    `xorm:"VARCHAR(60)"`
	Area         string    `xorm:"VARCHAR(60)"`
	Type         int       `xorm:"SMALLINT(1)"`
	Lang         string    `xorm:"VARCHAR(10)"`
	CreateUser   string    `xorm:"not null VARCHAR(30)"`
	CreateTime   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
