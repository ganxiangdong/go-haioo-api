package model

type HaiouFeed struct {
	Id       int    `xorm:"not null pk autoincr INT(10)"`
	Userid   int    `xorm:"INT(10)"`
	Company  string `xorm:"VARCHAR(100)"`
	Contact  string `xorm:"VARCHAR(30)"`
	Email    string `xorm:"VARCHAR(30)"`
	Mes      string `xorm:"TEXT"`
	Iflook   string `xorm:"CHAR(2)"`
	Province string `xorm:"VARCHAR(30)"`
	City     string `xorm:"VARCHAR(30)"`
	Tell     string `xorm:"VARCHAR(30)"`
	Addr     string `xorm:"VARCHAR(100)"`
}
