package model

import (
	"time"
)

type HaiouStockPurchasesPaper struct {
	Id         int       `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Title      string    `xorm:"VARCHAR(200)"`
	Shopid     int       `xorm:"not null default 0 MEDIUMINT(9)"`
	Shop       string    `xorm:"VARCHAR(40)"`
	Supplierid int       `xorm:"not null default 0 MEDIUMINT(9)"`
	Supplier   string    `xorm:"VARCHAR(60)"`
	Memo       string    `xorm:"TEXT"`
	CreateTime time.Time `xorm:"not null default '0000-00-00 00:00:00' TIMESTAMP"`
	Puser      string    `xorm:"VARCHAR(100)"`
	Status     int       `xorm:"default 0 TINYINT(1)"`
}
