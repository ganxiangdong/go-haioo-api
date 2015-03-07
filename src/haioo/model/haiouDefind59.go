package model

type HaiouDefind59 struct {
	Id       int    `xorm:"not null pk autoincr INT(11)"`
	InfoId   int    `xorm:"index(info_id) INT(11)"`
	InfoType string `xorm:"index(info_id) VARCHAR(30)"`
	ColorImg string `xorm:"TEXT"`
	Color    string `xorm:"VARCHAR(255)"`
	S377     string `xorm:"VARCHAR(255)"`
	S379     string `xorm:"VARCHAR(255)"`
	S381     string `xorm:"VARCHAR(255)"`
	G55      string `xorm:"VARCHAR(255)"`
}
