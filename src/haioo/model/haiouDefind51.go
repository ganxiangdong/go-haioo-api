package model

type HaiouDefind51 struct {
	Id       int    `xorm:"not null pk autoincr INT(11)"`
	InfoId   int    `xorm:"index(info_id) INT(11)"`
	InfoType string `xorm:"index(info_id) VARCHAR(30)"`
	ColorImg string `xorm:"TEXT"`
	S353     string `xorm:"VARCHAR(255)"`
	S355     string `xorm:"VARCHAR(255)"`
	S357     string `xorm:"VARCHAR(255)"`
	S359     string `xorm:"VARCHAR(255)"`
	S361     string `xorm:"VARCHAR(255)"`
	S363     string `xorm:"VARCHAR(255)"`
	S365     string `xorm:"VARCHAR(255)"`
	Color    string `xorm:"VARCHAR(255)"`
}
