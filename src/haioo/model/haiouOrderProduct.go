package model

import (
	"time"
)

type HaiouOrderProduct struct {
	Id          int       `xorm:"not null pk autoincr INT(11)"`
	OrderId     string    `xorm:"index(order_id) VARCHAR(20)"`
	MemberId    int       `xorm:"not null index(order_id) INT(11)"`
	ProductId   int       `xorm:"INT(11)"`
	Pcatid      int       `xorm:"not null INT(11)"`
	ProductName string    `xorm:"not null VARCHAR(100)"`
	Property    string    `xorm:"VARCHAR(150)"`
	Pic         string    `xorm:"not null VARCHAR(500)"`
	Price       string    `xorm:"DECIMAL(10,2)"`
	Num         int       `xorm:"INT(5)"`
	IsComment   int       `xorm:"not null default 0 TINYINT(1)"`
	CreateTime  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
