package model

type HaiouDefind53 struct {
	Id       int    `xorm:"not null pk autoincr INT(11)"`
	InfoId   int    `xorm:"index(info_id) INT(11)"`
	InfoType string `xorm:"index(info_id) VARCHAR(30)"`
	ColorImg string `xorm:"TEXT"`
	S277     string `xorm:"VARCHAR(255)"`
	S279     string `xorm:"VARCHAR(255)"`
	S269     string `xorm:"VARCHAR(255)"`
	S271     string `xorm:"VARCHAR(255)"`
	S273     string `xorm:"VARCHAR(255)"`
	S275     string `xorm:"VARCHAR(255)"`
	S267     string `xorm:"VARCHAR(255)"`
	Color    string `xorm:"VARCHAR(255)"`
	G39      string `xorm:"VARCHAR(255)"`
	S387     string `xorm:"VARCHAR(255)"`
	S389     string `xorm:"VARCHAR(255)"`
	S391     string `xorm:"VARCHAR(255)"`
}
