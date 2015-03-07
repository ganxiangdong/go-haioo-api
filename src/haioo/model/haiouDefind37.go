package model

type HaiouDefind37 struct {
	Id       int    `xorm:"not null pk autoincr INT(11)"`
	InfoId   int    `xorm:"index(info_id) INT(11)"`
	InfoType string `xorm:"index(info_id) VARCHAR(30)"`
	ColorImg string `xorm:"TEXT"`
	S285     string `xorm:"VARCHAR(255)"`
	S287     string `xorm:"VARCHAR(255)"`
	G49      string `xorm:"VARCHAR(255)"`
}
