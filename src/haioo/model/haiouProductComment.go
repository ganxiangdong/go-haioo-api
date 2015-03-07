package model

import (
	"time"
)

type HaiouProductComment struct {
	Id             int       `xorm:"not null pk autoincr INT(11)"`
	MemberId       int       `xorm:"not null index(member_id) INT(11)"`
	Username       string    `xorm:"not null CHAR(20)"`
	OrderId        string    `xorm:"not null index(member_id) VARCHAR(20)"`
	ProductSkuId   int       `xorm:"not null index(member_id) INT(10)"`
	ProductId      int       `xorm:"not null index(member_id) INT(11)"`
	ProductName    string    `xorm:"not null VARCHAR(100)"`
	Con            string    `xorm:"not null TEXT"`
	Pic            string    `xorm:"VARCHAR(500)"`
	Goodbad        int       `xorm:"not null default 0 TINYINT(1)"`
	ClientSystem   int       `xorm:"TINYINT(1)"`
	ClientPlatform int       `xorm:"TINYINT(1)"`
	CreateTime     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
