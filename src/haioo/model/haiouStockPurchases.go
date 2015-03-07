package model

type HaiouStockPurchases struct {
	Id          int     `xorm:"not null pk autoincr MEDIUMINT(8)"`
	PaperId     int     `xorm:"not null MEDIUMINT(11)"`
	Shopid      int     `xorm:"not null default 0 MEDIUMINT(9)"`
	Shop        string  `xorm:"VARCHAR(40)"`
	Productid   int     `xorm:"not null default 0 MEDIUMINT(9)"`
	Lotnumber   string  `xorm:"VARCHAR(90)"`
	Product     string  `xorm:"VARCHAR(60)"`
	Spec        string  `xorm:"VARCHAR(20)"`
	Unit        string  `xorm:"VARCHAR(10)"`
	Pricebuy    float32 `xorm:"not null default 0.00 FLOAT(10,2)"`
	Number      float32 `xorm:"not null default 0.00 FLOAT(10,2)"`
	Delnumber   float32 `xorm:"not null default 0.00 FLOAT(10,2)"`
	Type        int     `xorm:"not null default 0 TINYINT(1)"`
	Supplierid  int     `xorm:"not null default 0 MEDIUMINT(9)"`
	Supplier    string  `xorm:"VARCHAR(60)"`
	Memo        string  `xorm:"TEXT"`
	Timestart   int     `xorm:"INT(11)"`
	Usedby      string  `xorm:"VARCHAR(20)"`
	Timeend     int     `xorm:"INT(11)"`
	Thedate     int     `xorm:"not null default 0 INT(11)"`
	Timestocked int     `xorm:"default 0 INT(11)"`
	Puser       string  `xorm:"VARCHAR(100)"`
	Suser       string  `xorm:"VARCHAR(100)"`
	Timeexpress int     `xorm:"INT(11)"`
}
