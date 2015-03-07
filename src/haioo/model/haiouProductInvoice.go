package model

import (
	"time"
)

type HaiouProductInvoice struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	MemberId   int       `xorm:"not null INT(11)"`
	Type       int       `xorm:"not null default 0 TINYINT(1)"`
	Rise       int       `xorm:"not null default 0 TINYINT(1)"`
	Content    int       `xorm:"not null default 0 TINYINT(1)"`
	Company    string    `xorm:"VARCHAR(50)"`
	Number     string    `xorm:"VARCHAR(30)"`
	Address    string    `xorm:"VARCHAR(30)"`
	Telephone  string    `xorm:"VARCHAR(30)"`
	Bank       string    `xorm:"VARCHAR(30)"`
	Account    string    `xorm:"VARCHAR(20)"`
	IsDefault  int       `xorm:"TINYINT(1)"`
	Checked    int       `xorm:"not null default 0 TINYINT(1)"`
	CreateTime time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
