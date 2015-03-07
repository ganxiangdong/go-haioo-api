package model

import (
	"time"
)

type HaiouCaptcha struct {
	Id          int       `xorm:"not null pk autoincr INT(10)"`
	UniqueStr   string    `xorm:"not null index(unique_str) VARCHAR(32)"`
	Type        string    `xorm:"not null index(unique_str) VARCHAR(20)"`
	Captcha     string    `xorm:"not null index(unique_str) CHAR(6)"`
	IsUsed      int       `xorm:"not null default 0 index(unique_str) TINYINT(1)"`
	InvalidTime time.Time `xorm:"not null index(unique_str) DATETIME"`
	CreateTime  time.Time `xorm:"not null DATETIME"`
}
