package model

import (
	"time"
)

type HaiouPropertyValueTemplate struct {
	Id          int       `xorm:"not null pk autoincr INT(6)"`
	Field       string    `xorm:"VARCHAR(100)"`
	FieldName   string    `xorm:"VARCHAR(200)"`
	FieldDesc   string    `xorm:"VARCHAR(200)"`
	DisplayType int       `xorm:"INT(1)"`
	Item        string    `xorm:"VARCHAR(20000)"`
	IsSearch    int       `xorm:"default 0 TINYINT(1)"`
	Status      int       `xorm:"default 1 TINYINT(1)"`
	Type        int       `xorm:"default 1 TINYINT(1)"`
	CreateUser  string    `xorm:"VARCHAR(30)"`
	CreateTime  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
