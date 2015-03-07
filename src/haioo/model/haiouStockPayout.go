package model

type HaiouStockPayout struct {
	Id          int `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Ordersid    int `xorm:"not null default 0 MEDIUMINT(8)"`
	Purchasesid int `xorm:"not null MEDIUMINT(8)"`
	Number      int `xorm:"not null INT(11)"`
	Shopid      int `xorm:"not null default 0 MEDIUMINT(8)"`
	Productid   int `xorm:"not null default 0 MEDIUMINT(8)"`
}
