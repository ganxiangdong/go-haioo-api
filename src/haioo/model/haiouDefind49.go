package model

type HaiouDefind49 struct {
	Id       int    `xorm:"not null pk autoincr INT(11)"`
	InfoId   int    `xorm:"index(info_id) INT(11)"`
	InfoType string `xorm:"index(info_id) VARCHAR(30)"`
	ColorImg string `xorm:"TEXT"`
	S333     string `xorm:"VARCHAR(255)"`
	S335     string `xorm:"VARCHAR(255)"`
	S337     string `xorm:"VARCHAR(255)"`
	S339     string `xorm:"VARCHAR(255)"`
}
