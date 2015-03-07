package model

import (
	"time"
)

type HaiouPaymentCashflow struct {
	Id         int       `xorm:"not null pk autoincr INT(10)"`
	MemberId   int       `xorm:"index INT(11)"`
	Price      string    `xorm:"DECIMAL(10,2)"`
	Note       string    `xorm:"VARCHAR(255)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
