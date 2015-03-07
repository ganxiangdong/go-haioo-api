package model

type HaiouWebConfig struct {
	Id     int    `xorm:"not null pk autoincr INT(5)"`
	Index  string `xorm:"not null index(index) VARCHAR(30)"`
	Value  string `xorm:"not null TEXT"`
	Status int    `xorm:"not null default 1 TINYINT(1)"`
	Type   string `xorm:"index(index) VARCHAR(50)"`
	Des    string `xorm:"TEXT"`
}
