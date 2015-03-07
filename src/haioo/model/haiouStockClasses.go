package model

type HaiouStockClasses struct {
	Id             int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Title          string `xorm:"VARCHAR(20)"`
	Parentid       int    `xorm:"not null default 0 MEDIUMINT(9)"`
	Path           string `xorm:"VARCHAR(20)"`
	Ordersn        int    `xorm:"not null default 0 MEDIUMINT(9)"`
	Sequence       int    `xorm:"not null default 0 MEDIUMINT(9)"`
	Priceline      int    `xorm:"TINYINT(2)"`
	Priceunderline int    `xorm:"TINYINT(2)"`
	Warningnum     int    `xorm:"INT(10)"`
	Canbuyout      int    `xorm:"INT(10)"`
}
