package model

type HaiouDefind41 struct {
	Id       int    `xorm:"not null pk autoincr INT(11)"`
	InfoId   int    `xorm:"index(info_id) INT(11)"`
	InfoType string `xorm:"index(info_id) VARCHAR(30)"`
	ColorImg string `xorm:"TEXT"`
	S291     string `xorm:"VARCHAR(255)"`
	S293     string `xorm:"VARCHAR(255)"`
	S295     string `xorm:"VARCHAR(255)"`
}
