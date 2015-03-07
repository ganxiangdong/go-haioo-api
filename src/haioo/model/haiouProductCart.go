package model

import (
	"time"
)

type HaiouProductCart struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	Phoneid    string    `xorm:"default '0' index(phoneid) VARCHAR(30)"`
	MemberId   int       `xorm:"index(member_id) INT(11)"`
	ProductId  int       `xorm:"not null index(phoneid) index(member_id) INT(11)"`
	Price      string    `xorm:"not null DECIMAL(10,2)"`
	Weight     float64   `xorm:"not null DOUBLE(10,3)"`
	Num        int       `xorm:"not null INT(5)"`
	Status     int       `xorm:"not null default 1 TINYINT(1)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateTime int       `xorm:"not null INT(11)"`
}
