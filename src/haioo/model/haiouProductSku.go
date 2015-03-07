package model

import (
	"time"
)

type HaiouProductSku struct {
	Id              int       `xorm:"not null pk autoincr INT(10)"`
	Pid             int       `xorm:"not null default 0 index INT(10)"`
	Catid           int       `xorm:"INT(8)"`
	Code            string    `xorm:"default '0' VARCHAR(60)"`
	Pname           string    `xorm:"not null VARCHAR(100)"`
	Spec            string    `xorm:"VARCHAR(255)"`
	PromotionTips   string    `xorm:"VARCHAR(255)"`
	Price           string    `xorm:"not null default 0.00 index DECIMAL(10,2)"`
	Unit            int       `xorm:"not null default 1 TINYINT(2)"`
	ShowPrice       string    `xorm:"not null DECIMAL(10,2)"`
	MarketPrice     string    `xorm:"default 0.00 DECIMAL(10,2)"`
	TaxRate         float64   `xorm:"DOUBLE(8,3)"`
	CostPrice       string    `xorm:"default 0.00 DECIMAL(10,2)"`
	StockRelations  string    `xorm:"VARCHAR(255)"`
	Stock           int       `xorm:"not null default 0 INT(11)"`
	Tags            string    `xorm:"VARCHAR(255)"`
	Brandid         int       `xorm:"INT(6)"`
	Brand           string    `xorm:"VARCHAR(60)"`
	Pic             string    `xorm:"VARCHAR(1500)"`
	SellAmount      int       `xorm:"default 0 INT(6)"`
	Province        int       `xorm:"INT(11)"`
	City            int       `xorm:"INT(6)"`
	Areaid          int       `xorm:"INT(6)"`
	Area            string    `xorm:"VARCHAR(255)"`
	Uptime          int       `xorm:"INT(10)"`
	Status          int       `xorm:"default 0 TINYINT(1)"`
	Weight          float64   `xorm:"DOUBLE(10,3)"`
	Cubage          float64   `xorm:"DOUBLE(10,3)"`
	PropertyValueId string    `xorm:"TEXT"`
	CreateUser      string    `xorm:"VARCHAR(30)"`
	CreateTime      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
