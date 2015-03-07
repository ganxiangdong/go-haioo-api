package model

import (
	"time"
)

type HaiouMemberAddress struct {
	Id               int       `xorm:"not null pk autoincr INT(11)"`
	Userid           int       `xorm:"not null index(userid) INT(11)"`
	Name             string    `xorm:"not null VARCHAR(40)"`
	Provinceid       int       `xorm:"not null INT(11)"`
	Cityid           int       `xorm:"not null INT(11)"`
	Areaid           int       `xorm:"not null INT(11)"`
	Area             string    `xorm:"not null VARCHAR(255)"`
	Address          string    `xorm:"not null VARCHAR(50)"`
	Zip              int       `xorm:"INT(10)"`
	Tel              string    `xorm:"VARCHAR(30)"`
	Mobile           string    `xorm:"VARCHAR(20)"`
	Email            string    `xorm:"VARCHAR(255)"`
	Idcard           string    `xorm:"VARCHAR(18)"`
	AddressLatitude  string    `xorm:"VARCHAR(10)"`
	AddressLongitude string    `xorm:"VARCHAR(10)"`
	IsDefault        int       `xorm:"not null default 0 index(userid) TINYINT(4)"`
	Status           int       `xorm:"not null index(userid) TINYINT(1)"`
	CreateTime       time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
