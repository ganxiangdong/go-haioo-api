package model

import (
	"time"
)

type HaiouAppVersion struct {
	Id              int       `xorm:"not null pk autoincr INT(10)"`
	Version         float32   `xorm:"not null index FLOAT(3,1)"`
	ClientSystem    int       `xorm:"not null TINYINT(1)"`
	DownloadUrl     string    `xorm:"not null VARCHAR(180)"`
	Description     string    `xorm:"not null VARCHAR(500)"`
	MustUpdatedRule string    `xorm:"VARCHAR(120)"`
	CreateTime      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
