package model

import (
	"time"
)

type HaiouActivityFavorite struct {
	Id                int       `xorm:"not null pk autoincr INT(10)"`
	MemberId          int       `xorm:"not null index(activity_id) INT(10)"`
	ActivityId        int       `xorm:"not null index(activity_id) INT(10)"`
	ActivityStartTime int       `xorm:"not null INT(10)"`
	ClientPlatform    int       `xorm:"TINYINT(1)"`
	ClientSystem      int       `xorm:"TINYINT(1)"`
	CreateTime        time.Time `xorm:"DATETIME"`
}
