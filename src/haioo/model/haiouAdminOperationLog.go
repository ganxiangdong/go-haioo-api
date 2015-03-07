package model

type HaiouAdminOperationLog struct {
	Id         int    `xorm:"not null pk autoincr INT(5)"`
	User       string `xorm:"VARCHAR(20)"`
	Scriptname string `xorm:"VARCHAR(50)"`
	Url        string `xorm:"VARCHAR(200)"`
	Time       int    `xorm:"INT(11)"`
}
