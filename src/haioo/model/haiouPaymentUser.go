package model

import (
	"time"
)

type HaiouPaymentUser struct {
	MemberId    int       `xorm:"not null pk default 0 INT(11)"`
	PayPass     string    `xorm:"VARCHAR(32)"`
	Cash        string    `xorm:"default 0.00 DECIMAL(10,2)"`
	Unreachable string    `xorm:"default 0.00 DECIMAL(10,2)"`
	CreateTime  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
