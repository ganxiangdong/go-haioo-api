package model

import (
	"time"
)

type HaiouLogisticsExpress struct {
	Id           int       `xorm:"not null pk autoincr INT(11)"`
	Company      string    `xorm:"VARCHAR(30)"`
	Introduction string    `xorm:"TEXT"`
	Url          string    `xorm:"VARCHAR(30)"`
	Logo         string    `xorm:"VARCHAR(300)"`
	Sign         string    `xorm:"VARCHAR(10)"`
	CreateUser   string    `xorm:"VARCHAR(30)"`
	CreateTime   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Html         string    `xorm:"TEXT"`
	Regular      string    `xorm:"VARCHAR(500)"`
}
