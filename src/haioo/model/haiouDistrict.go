package model

type HaiouDistrict struct {
	Id      int    `xorm:"not null pk autoincr INT(11)"`
	Name    string `xorm:"not null index VARCHAR(15)"`
	Pid     int    `xorm:"not null index INT(11)"`
	Zip     int    `xorm:"INT(6)"`
	Sorting int    `xorm:"TINYINT(3)"`
}
