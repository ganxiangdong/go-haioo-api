package model

import (
	"time"
)

type HaiouAdminGroup struct {
	GroupId    int       `xorm:"not null pk autoincr SMALLINT(5)"`
	GroupName  string    `xorm:"not null VARCHAR(60)"`
	GroupPerms string    `xorm:"not null TEXT"`
	GroupDesc  string    `xorm:"not null VARCHAR(250)"`
	CreateUser string    `xorm:"not null VARCHAR(30)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
