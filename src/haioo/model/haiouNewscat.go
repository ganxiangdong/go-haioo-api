package model

type HaiouNewscat struct {
	Catid    int    `xorm:"not null pk autoincr INT(6)"`
	Cat      string `xorm:"CHAR(100)"`
	Nums     int    `xorm:"INT(10)"`
	Pid      int    `xorm:"default 0 INT(6)"`
	Ishome   int    `xorm:"INT(1)"`
	Template string `xorm:"VARCHAR(50)"`
	Pic      string `xorm:"VARCHAR(100)"`
	Openpost int    `xorm:"default 0 TINYINT(1)"`
	Desc     string `xorm:"VARCHAR(255)"`
}
