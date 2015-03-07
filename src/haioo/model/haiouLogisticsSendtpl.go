package model

type HaiouLogisticsSendtpl struct {
	Id      int    `xorm:"not null pk autoincr INT(11)"`
	Title   string `xorm:"VARCHAR(255)"`
	Type    int    `xorm:"TINYINT(1)"`
	Default int    `xorm:"default 0 TINYINT(1)"`
	Sendtpl string `xorm:"TEXT"`
	Status  int    `xorm:"default 1 TINYINT(1)"`
}
