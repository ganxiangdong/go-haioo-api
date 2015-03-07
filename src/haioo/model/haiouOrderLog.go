package model

import (
	"time"
)

type HaiouOrderLog struct {
	Id         int       `xorm:"not null pk autoincr INT(20)"`
	OrderId    string    `xorm:"VARCHAR(20)"`
	AdminId    int       `xorm:"SMALLINT(5)"`
	AdminName  string    `xorm:"VARCHAR(30)"`
	Text       string    `xorm:"LONGTEXT"`
	Behavior   string    `xorm:"default '' VARCHAR(20)"`
	Result     string    `xorm:"default 'success' ENUM('success','failure')"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
