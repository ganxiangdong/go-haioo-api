package model

type HaiouMailMod struct {
	Id      int    `xorm:"not null pk autoincr INT(8)"`
	Subject string `xorm:"VARCHAR(100)"`
	Title   string `xorm:"VARCHAR(100)"`
	Message string `xorm:"TEXT"`
	Flag    string `xorm:"index(flag) VARCHAR(30)"`
	Type    int    `xorm:"not null default 0 index(flag) TINYINT(1)"`
}
