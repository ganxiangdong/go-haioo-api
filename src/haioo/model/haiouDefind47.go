package model

type HaiouDefind47 struct {
	Id       int    `xorm:"not null pk autoincr INT(11)"`
	InfoId   int    `xorm:"index(info_id) INT(11)"`
	InfoType string `xorm:"index(info_id) VARCHAR(30)"`
	ColorImg string `xorm:"TEXT"`
	S341     string `xorm:"VARCHAR(255)"`
	S343     string `xorm:"VARCHAR(255)"`
	S345     string `xorm:"VARCHAR(255)"`
	S347     string `xorm:"VARCHAR(255)"`
	Color    string `xorm:"VARCHAR(255)"`
	G39      string `xorm:"VARCHAR(255)"`
}
