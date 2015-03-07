package model

import (
	"time"
)

type HaiouAdvsCon struct {
	Id         int       `xorm:"not null pk autoincr INT(4)"`
	Userid     int       `xorm:"INT(11)"`
	GroupId    int       `xorm:"index(group_id) INT(5)"`
	Name       string    `xorm:"not null default '' VARCHAR(50)"`
	Type       string    `xorm:"VARCHAR(20)"`
	Url        string    `xorm:"VARCHAR(200)"`
	Con        string    `xorm:"MEDIUMTEXT"`
	Isopen     int       `xorm:"default 0 index(group_id) INT(1)"`
	Ctime      int       `xorm:"INT(11)"`
	Province   string    `xorm:"VARCHAR(50)"`
	City       string    `xorm:"VARCHAR(50)"`
	Area       string    `xorm:"VARCHAR(50)"`
	Pic        string    `xorm:"VARCHAR(255)"`
	Width      string    `xorm:"CHAR(4)"`
	Height     string    `xorm:"CHAR(4)"`
	Catid      int       `xorm:"INT(8)"`
	Unit       string    `xorm:"ENUM('day','week','month')"`
	ShowTime   int       `xorm:"default 0 TINYINT(4)"`
	Status     int       `xorm:"default 0 TINYINT(1)"`
	Shownum    int       `xorm:"default 1 INT(11)"`
	Stime      int       `xorm:"index(group_id) INT(10)"`
	Etime      int       `xorm:"index(group_id) INT(10)"`
	SortNum    int       `xorm:"default 0 TINYINT(3)"`
	CreateUser string    `xorm:"VARCHAR(30)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	CatType    int       `xorm:"not null default 9 TINYINT(1)"`
}
