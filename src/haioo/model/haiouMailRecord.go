package model

import (
	"time"
)

type HaiouMailRecord struct {
	Id             int       `xorm:"not null pk autoincr INT(11)"`
	Sendmailname   string    `xorm:"VARCHAR(20)"`
	Sendtime       time.Time `xorm:"DATETIME"`
	Sendmailrecord string    `xorm:"TEXT"`
}
