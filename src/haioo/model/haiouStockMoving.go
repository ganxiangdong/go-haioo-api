package model

type HaiouStockMoving struct {
	Id          int     `xorm:"not null pk autoincr MEDIUMINT(8)"`
	ShopidFrom  int     `xorm:"not null default 0 MEDIUMINT(9)"`
	ShopFrom    string  `xorm:"VARCHAR(40)"`
	ShopidTo    int     `xorm:"not null default 0 MEDIUMINT(9)"`
	ShopTo      string  `xorm:"VARCHAR(40)"`
	Productid   int     `xorm:"not null default 0 MEDIUMINT(9)"`
	Product     string  `xorm:"VARCHAR(60)"`
	Spec        string  `xorm:"VARCHAR(20)"`
	Unit        string  `xorm:"VARCHAR(10)"`
	Number      float32 `xorm:"not null default 0.00 FLOAT(10,2)"`
	Memo        string  `xorm:"TEXT"`
	Thedate     int     `xorm:"not null default 0 INT(11)"`
	Timestocked int     `xorm:"not null default 0 INT(11)"`
}
