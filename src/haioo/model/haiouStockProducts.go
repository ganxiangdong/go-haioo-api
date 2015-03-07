package model

type HaiouStockProducts struct {
	Id             int     `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Classid        int     `xorm:"default 0 MEDIUMINT(9)"`
	Path           string  `xorm:"VARCHAR(20)"`
	Pricebuy       float32 `xorm:"FLOAT(10,2)"`
	Pricesell      float32 `xorm:"FLOAT(10,2)"`
	Priceundersell float32 `xorm:"FLOAT(10,2)"`
	Unit           string  `xorm:"VARCHAR(10)"`
	Intro          string  `xorm:"TEXT"`
	Stocks         float32 `xorm:"default 0.00 FLOAT(10,2)"`
	Maxnum         int     `xorm:"default 1 INT(11)"`
	Minnum         int     `xorm:"default 1 INT(11)"`
	Shopwarningnum int     `xorm:"not null default 5 INT(5)"`
	Canbuyout      int     `xorm:"default 0 INT(11)"`
	ProductSkuId   int     `xorm:"not null INT(10)"`
}
