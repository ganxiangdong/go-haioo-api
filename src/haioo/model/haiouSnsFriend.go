package model

import (
	"time"
)

type HaiouSnsFriend struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	MemberId   int       `xorm:"not null INT(11)"`
	MemberName string    `xorm:"VARCHAR(100)"`
	MemberImg  string    `xorm:"VARCHAR(100)"`
	FriendId   int       `xorm:"not null INT(11)"`
	FriendName string    `xorm:"not null VARCHAR(100)"`
	FriendImg  string    `xorm:"VARCHAR(100)"`
	Status     int       `xorm:"not null default 1 TINYINT(1)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
