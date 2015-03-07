package model

type HaiouDefind39 struct {
	Id       int    `xorm:"not null pk autoincr INT(11)"`
	InfoId   int    `xorm:"index(info_id) INT(11)"`
	InfoType string `xorm:"index(info_id) VARCHAR(30)"`
	ColorImg string `xorm:"TEXT"`
	S309     string `xorm:"VARCHAR(255)"`
	S311     string `xorm:"VARCHAR(255)"`
	S313     string `xorm:"VARCHAR(255)"`
	S315     string `xorm:"VARCHAR(255)"`
	S317     string `xorm:"VARCHAR(255)"`
	G49      string `xorm:"VARCHAR(255)"`
}
