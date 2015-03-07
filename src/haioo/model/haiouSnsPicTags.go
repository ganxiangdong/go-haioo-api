package model

import (
	"time"
)

type HaiouSnsPicTags struct {
	Id         int       `xorm:"not null pk autoincr INT(10)"`
	MemberId   int       `xorm:"not null index(member_id) INT(11)"`
	SnsId      int       `xorm:"not null INT(10)"`
	Tagname    string    `xorm:"not null index(member_id) VARCHAR(100)"`
	Brandid    int       `xorm:"not null index INT(10)"`
	Brand      string    `xorm:"VARCHAR(80)"`
	Type       int       `xorm:"not null default 1 TINYINT(1)"`
	PointX     float64   `xorm:"DOUBLE(8,3)"`
	PointY     float64   `xorm:"DOUBLE(8,3)"`
	Direction  int       `xorm:"default 4 TINYINT(1)"`
	Status     int       `xorm:"not null TINYINT(1)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
