package model

type HaiouDefind43 struct {
	Id       int    `xorm:"not null pk autoincr INT(11)"`
	InfoId   int    `xorm:"index(info_id) INT(11)"`
	InfoType string `xorm:"index(info_id) VARCHAR(30)"`
	ColorImg string `xorm:"TEXT"`
	S305     string `xorm:"VARCHAR(255)"`
	S303     string `xorm:"VARCHAR(255)"`
	S301     string `xorm:"VARCHAR(255)"`
	S299     string `xorm:"VARCHAR(255)"`
	S297     string `xorm:"VARCHAR(255)"`
	G49      string `xorm:"VARCHAR(255)"`
}
