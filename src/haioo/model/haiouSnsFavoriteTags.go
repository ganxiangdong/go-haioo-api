package model

import (
	"time"
)

type HaiouSnsFavoriteTags struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	MemberId   int       `xorm:"not null index(member_id) INT(11)"`
	TagId      int       `xorm:"not null index(member_id) index INT(11)"`
	Tagname    string    `xorm:"VARCHAR(100)"`
	CreateTime time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
