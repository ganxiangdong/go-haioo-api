package model

import (
	"time"
)

type HaiouOrder struct {
	Id                  int       `xorm:"not null pk autoincr INT(20)"`
	MemberId            int       `xorm:"unique(order_id) INT(11)"`
	OrderId             string    `xorm:"unique(order_id) VARCHAR(20)"`
	SellerId            int       `xorm:"INT(11)"`
	Status              int       `xorm:"default 1 unique(order_id) TINYINT(1)"`
	PayStatus           int       `xorm:"default 1 TINYINT(1)"`
	ShipStatus          int       `xorm:"default 1 TINYINT(1)"`
	OrderTime           int       `xorm:"INT(10)"`
	ShipTime            int       `xorm:"INT(10)"`
	PayTime             int       `xorm:"INT(10)"`
	FinishedTime        int       `xorm:"INT(10)"`
	InvoiceId           int       `xorm:"default 0 INT(11)"`
	PaymentId           int       `xorm:"default 0 SMALLINT(4)"`
	PaymentName         string    `xorm:"VARCHAR(100)"`
	ShippingId          int       `xorm:"default 0 MEDIUMINT(8)"`
	ShippingName        string    `xorm:"VARCHAR(100)"`
	ShipName            string    `xorm:"VARCHAR(50)"`
	ShipAddr            string    `xorm:"VARCHAR(255)"`
	ShipTel             string    `xorm:"VARCHAR(30)"`
	ShipMobile          string    `xorm:"VARCHAR(15)"`
	ShipZip             string    `xorm:"VARCHAR(6)"`
	ShippingIdcard      string    `xorm:"VARCHAR(18)"`
	ShippingExpressId   int       `xorm:"MEDIUMINT(8)"`
	ShippingExpressName string    `xorm:"VARCHAR(100)"`
	ShippingNo          string    `xorm:"VARCHAR(50)"`
	Des                 string    `xorm:"VARCHAR(200)"`
	Cost                string    `xorm:"default 0.00 DECIMAL(20,2)"`
	CpnsPrice           string    `xorm:"default 0.00 DECIMAL(10,2)"`
	CashPrice           string    `xorm:"default 0.00 DECIMAL(10,2)"`
	TaxPrice            string    `xorm:"default 0.00 DECIMAL(10,2)"`
	Freight             string    `xorm:"default 0.00 DECIMAL(20,2)"`
	Weight              string    `xorm:"default 0.00 DECIMAL(20,2)"`
	IsComment           int       `xorm:"default 0 TINYINT(1)"`
	AdminRemark         string    `xorm:"VARCHAR(255)"`
	ClientPlatform      int       `xorm:"TINYINT(1)"`
	ClientSystem        int       `xorm:"TINYINT(1)"`
	CreateTime          time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
