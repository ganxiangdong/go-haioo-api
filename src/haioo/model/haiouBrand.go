package model

import (
	"time"
)

type HaiouBrand struct {
	Id           int       `xorm:"not null pk autoincr INT(11)"`
	Name         string    `xorm:"not null VARCHAR(80)"`
	SeoName      string    `xorm:"default '' VARCHAR(80)"`
	CharIndex    string    `xorm:"not null CHAR(1)"`
	Catid        int       `xorm:"not null INT(11)"`
	Catname      string    `xorm:"not null VARCHAR(20)"`
	Logo         string    `xorm:"not null VARCHAR(150)"`
	LogoLarge    string    `xorm:"default '' VARCHAR(150)"`
	Displayorder int       `xorm:"not null default 0 SMALLINT(6)"`
	Pic          string    `xorm:"VARCHAR(1500)"`
	Story        string    `xorm:"TEXT"`
	Status       int       `xorm:"not null default 1 TINYINT(1)"`
	CreateUser   string    `xorm:"VARCHAR(30)"`
	CreateTime   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
