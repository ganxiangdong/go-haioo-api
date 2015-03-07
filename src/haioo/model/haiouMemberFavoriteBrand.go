package model

import (
	"time"
)

type HaiouMemberFavoriteBrand struct {
	Id             int       `xorm:"not null pk autoincr INT(10)"`
	MemberId       int       `xorm:"not null index(member_id) INT(10)"`
	ActivityId     int       `xorm:"not null index(member_id) INT(10)"`
	ActivityTitle  string    `xorm:"VARCHAR(100)"`
	Brandid        int       `xorm:"not null INT(10)"`
	Brand          string    `xorm:"not null VARCHAR(60)"`
	Status         int       `xorm:"not null index(member_id) TINYINT(1)"`
	Displayorder   int       `xorm:"not null default 0 SMALLINT(6)"`
	ClientPlatform int       `xorm:"TINYINT(1)"`
	ClientSystem   int       `xorm:"TINYINT(1)"`
	CreateTime     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
