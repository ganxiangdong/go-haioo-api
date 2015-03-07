package model

type HaiouWebCon struct {
	ConId       int    `xorm:"not null pk autoincr INT(5)"`
	ConDesc     string `xorm:"MEDIUMTEXT"`
	ConProvince string `xorm:"VARCHAR(20)"`
	ConCity     string `xorm:"VARCHAR(20)"`
	ConNo       int    `xorm:"default 0 INT(2)"`
	ConStatus   int    `xorm:"default 0 INT(1)"`
	ConTitle    string `xorm:"VARCHAR(30)"`
	ConLinkaddr string `xorm:"VARCHAR(60)"`
	ConGroup    int    `xorm:"not null TINYINT(3)"`
	Template    string `xorm:"VARCHAR(50)"`
	Title       string `xorm:"VARCHAR(200)"`
	Keywords    string `xorm:"VARCHAR(200)"`
	Description string `xorm:"VARCHAR(200)"`
	MsgOnline   int    `xorm:"not null default 0 TINYINT(1)"`
	Lang        string `xorm:"VARCHAR(5)"`
	Type        int    `xorm:"default 0 TINYINT(1)"`
}
