package model

import (
	"time"
)

type HaiouBrandCat struct {
	Id           int       `xorm:"not null pk autoincr INT(11)"`
	ParentId     int       `xorm:"not null default 0 index(parent_id) INT(11)"`
	Displayorder int       `xorm:"not null default 0 index(parent_id) SMALLINT(6)"`
	Catname      string    `xorm:"not null VARCHAR(100)"`
	ShortCode    string    `xorm:"not null CHAR(2)"`
	CreateUser   string    `xorm:"VARCHAR(30)"`
	CreateTime   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
