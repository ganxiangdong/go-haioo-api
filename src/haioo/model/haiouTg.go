package model

type HaiouTg struct {
	Id              int     `xorm:"not null pk autoincr INT(8)"`
	Catid           int     `xorm:"not null INT(6)"`
	Name            string  `xorm:"not null VARCHAR(50)"`
	Content         string  `xorm:"not null TEXT"`
	Pic             string  `xorm:"not null VARCHAR(255)"`
	MarketPrice     float32 `xorm:"not null default 0.00 FLOAT(10,2)"`
	Price           float32 `xorm:"not null default 0.00 FLOAT(10,2)"`
	Hits            int     `xorm:"not null default 0 INT(6)"`
	SellAmount      int     `xorm:"not null default 0 INT(6)"`
	LimitQuantity   int     `xorm:"not null default 0 INT(6)"`
	VirtualQuantity int     `xorm:"not null default 0 INT(6)"`
	Status          int     `xorm:"not null default 0 TINYINT(1)"`
	CreateTime      int     `xorm:"not null INT(10)"`
	Stock           int     `xorm:"not null default 1 INT(10)"`
	Provinceid      int     `xorm:"not null INT(11)"`
	Cityid          int     `xorm:"not null INT(11)"`
	Areaid          int     `xorm:"not null INT(11)"`
	Area            string  `xorm:"not null VARCHAR(255)"`
	Displayorder    int     `xorm:"not null default 255 SMALLINT(6)"`
}
