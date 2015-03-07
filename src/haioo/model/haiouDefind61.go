package model

type HaiouDefind61 struct {
	Id       int    `xorm:"not null pk autoincr INT(11)"`
	InfoId   int    `xorm:"index(info_id) INT(11)"`
	InfoType string `xorm:"index(info_id) VARCHAR(30)"`
	ColorImg string `xorm:"TEXT"`
	Color    string `xorm:"VARCHAR(255)"`
	S369     string `xorm:"VARCHAR(255)"`
	S371     string `xorm:"VARCHAR(255)"`
	S373     string `xorm:"VARCHAR(255)"`
}
