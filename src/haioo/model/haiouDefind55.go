package model

type HaiouDefind55 struct {
	Id       int    `xorm:"not null pk autoincr INT(11)"`
	InfoId   int    `xorm:"index(info_id) INT(11)"`
	InfoType string `xorm:"index(info_id) VARCHAR(30)"`
	ColorImg string `xorm:"TEXT"`
	Color    string `xorm:"VARCHAR(255)"`
	G57      string `xorm:"VARCHAR(255)"`
	S393     string `xorm:"VARCHAR(255)"`
	S395     string `xorm:"VARCHAR(255)"`
	S397     string `xorm:"VARCHAR(255)"`
}
