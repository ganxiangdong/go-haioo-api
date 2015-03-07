package model

import (
	"time"
)

type HaiouProducts struct {
	Id             int       `xorm:"not null pk autoincr INT(6)"`
	Catid          int       `xorm:"index(catid) INT(9)"`
	Code           string    `xorm:"VARCHAR(60)"`
	Pname          string    `xorm:"index index(catid) VARCHAR(100)"`
	PromotionTips  string    `xorm:"VARCHAR(255)"`
	Keywords       string    `xorm:"index VARCHAR(100)"`
	Tags           string    `xorm:"VARCHAR(255)"`
	Brandid        int       `xorm:"INT(10)"`
	Brand          string    `xorm:"VARCHAR(60)"`
	Pic            string    `xorm:"VARCHAR(1500)"`
	Goodbad        int       `xorm:"INT(6)"`
	Comments       int       `xorm:"INT(6)"`
	Detail         string    `xorm:"TEXT"`
	Price          string    `xorm:"not null DECIMAL(10,2)"`
	ShowPrice      string    `xorm:"not null DECIMAL(10,2)"`
	MarketPrice    string    `xorm:"DECIMAL(10,2)"`
	AllStock       int       `xorm:"default 0 TINYINT(7)"`
	Unit           int       `xorm:"not null default 1 TINYINT(2)"`
	TaxRate        float64   `xorm:"DOUBLE(8,3)"`
	Province       int       `xorm:"INT(11)"`
	City           int       `xorm:"INT(6)"`
	Areaid         int       `xorm:"INT(6)"`
	Area           string    `xorm:"VARCHAR(255)"`
	Weight         float64   `xorm:"DOUBLE(10,3)"`
	Cubage         float64   `xorm:"DOUBLE(10,3)"`
	ProductRelated string    `xorm:"VARCHAR(255)"`
	ProductGive    string    `xorm:"VARCHAR(255)"`
	ProductParts   string    `xorm:"VARCHAR(255)"`
	Status         int       `xorm:"not null default 0 TINYINT(1)"`
	CreateUser     string    `xorm:"VARCHAR(30)"`
	CreateTime     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Hscode         string    `xorm:"default '' VARCHAR(30)"`
	PropertyId     int       `xorm:"not null INT(10)"`
	WaitingTime    string    `xorm:"default '' VARCHAR(30)"`
	WaitingTimeEnd string    `xorm:"default '' VARCHAR(30)"`
}
