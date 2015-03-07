package model

import (
	"time"
)

type HaiouSnsMemberFavorite struct {
	Id         int       `xorm:"not null pk autoincr INT(10)"`
	SnsId      int       `xorm:"not null index(mbmer_id) INT(10)"`
	MemberId   int       `xorm:"not null index(mbmer_id) INT(10)"`
	CreateTime time.Time `xorm:"not null DATETIME"`
}
