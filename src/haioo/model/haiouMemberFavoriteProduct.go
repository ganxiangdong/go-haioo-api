package model

import (
	"time"
)

type HaiouMemberFavoriteProduct struct {
	Id             int       `xorm:"not null pk autoincr INT(10)"`
	MemberId       int       `xorm:"not null index(member_id) INT(10)"`
	ProductId      int       `xorm:"not null index(member_id) INT(10)"`
	ProductName    string    `xorm:"not null VARCHAR(100)"`
	Status         int       `xorm:"not null index(member_id) TINYINT(1)"`
	Displayorder   int       `xorm:"not null default 0 SMALLINT(6)"`
	ClientPlatform int       `xorm:"TINYINT(1)"`
	ClientSystem   int       `xorm:"TINYINT(1)"`
	CreateTime     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
