package model

type HaiouStockStocks struct {
	Id        int     `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Shopid    int     `xorm:"not null default 0 MEDIUMINT(9)"`
	Shop      string  `xorm:"VARCHAR(40)"`
	Productid int     `xorm:"not null default 0 MEDIUMINT(9)"`
	Product   string  `xorm:"VARCHAR(60)"`
	Spec      string  `xorm:"VARCHAR(20)"`
	Unit      string  `xorm:"VARCHAR(10)"`
	Stocks    float32 `xorm:"not null default 0.00 FLOAT(10,2)"`
}
