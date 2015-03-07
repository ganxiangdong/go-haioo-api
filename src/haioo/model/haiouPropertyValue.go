package model

import (
	"time"
)

type HaiouPropertyValue struct {
	Id           int       `xorm:"not null pk autoincr INT(6)"`
	Field        string    `xorm:"VARCHAR(100)"`
	FieldName    string    `xorm:"VARCHAR(200)"`
	FieldDesc    string    `xorm:"VARCHAR(100)"`
	DisplayType  int       `xorm:"INT(1)"`
	Item         string    `xorm:"TEXT"`
	IsSearch     int       `xorm:"default 0 TINYINT(1)"`
	PropertyId   int       `xorm:"index INT(11)"`
	Status       int       `xorm:"default 1 TINYINT(1)"`
	Type         int       `xorm:"default 1 TINYINT(1)"`
	TemplateId   int       `xorm:"default 0 INT(6)"`
	Displayorder int       `xorm:"default 255 SMALLINT(6)"`
	CreateUser   string    `xorm:"VARCHAR(30)"`
	CreateTime   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
