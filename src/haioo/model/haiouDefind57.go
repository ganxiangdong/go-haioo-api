package model

type HaiouDefind57 struct {
	Id       int    `xorm:"not null pk autoincr INT(11)"`
	InfoId   int    `xorm:"index(info_id) INT(11)"`
	InfoType string `xorm:"index(info_id) VARCHAR(30)"`
	ColorImg string `xorm:"TEXT"`
	Color    string `xorm:"VARCHAR(255)"`
	G57      string `xorm:"VARCHAR(255)"`
	S403     string `xorm:"VARCHAR(255)"`
	S405     string `xorm:"VARCHAR(255)"`
	S407     string `xorm:"VARCHAR(255)"`
}
