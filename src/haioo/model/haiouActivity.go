package model

import (
//	"time"
)

type HaiouActivity struct {
	Id           int    `xorm:"not null pk autoincr INT(10)"`
	Title        string `xorm:"not null VARCHAR(100)"`
	Desc         string `xorm:"not null TEXT"`
	Pic          string `xorm:"not null VARCHAR(200)"`
	Discount     string `xorm:"VARCHAR(50)"`
	StartTime    int    `xorm:"not null index(start_time) INT(10)"`
	EndTime      int    `xorm:"not null index(start_time) INT(10)"`
	Catid        int    `xorm:"not null INT(10)"`
	Tags         string `xorm:"VARCHAR(255)"`
	Brandid      int    `xorm:"not null INT(10)"`
	CreateUser   string `xorm:"not null VARCHAR(30)"`
	CreateTime   string `xorm:"not null"`
	Status       int    `xorm:"not null default 0 index(start_time) TINYINT(1)"`
	Displayorder int    `xorm:"not null default 0 SMALLINT(6)"`
}
