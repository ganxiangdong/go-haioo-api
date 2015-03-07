package model

type HaiouDefind45 struct {
	Id       int    `xorm:"not null pk autoincr INT(11)"`
	InfoId   int    `xorm:"index(info_id) INT(11)"`
	InfoType string `xorm:"index(info_id) VARCHAR(30)"`
	ColorImg string `xorm:"TEXT"`
	S321     string `xorm:"VARCHAR(255)"`
	S323     string `xorm:"VARCHAR(255)"`
	S325     string `xorm:"VARCHAR(255)"`
	S327     string `xorm:"VARCHAR(255)"`
	S329     string `xorm:"VARCHAR(255)"`
	Color    string `xorm:"VARCHAR(255)"`
}
