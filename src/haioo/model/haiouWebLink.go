package model

type HaiouWebLink struct {
	Linkid   int    `xorm:"not null pk autoincr INT(4)"`
	Name     string `xorm:"VARCHAR(20)"`
	Url      string `xorm:"VARCHAR(200)"`
	Status   int    `xorm:"not null default 0 TINYINT(1)"`
	Orderid  int    `xorm:"not null default 0 INT(11)"`
	Log      string `xorm:"VARCHAR(100)"`
	Province string `xorm:"VARCHAR(15)"`
	City     string `xorm:"VARCHAR(15)"`
	Stime    int    `xorm:"INT(11)"`
	Etime    int    `xorm:"INT(11)"`
	Con      string `xorm:"TEXT"`
}
