package model

import (
	"time"
)

type HaiouActivityProductList struct {
	Id           int       `xorm:"not null pk autoincr INT(10)"`
	ActivityId   int       `xorm:"not null index(activity_id) INT(10)"`
	ProductId    int       `xorm:"not null INT(10)"`
	StartTime    int       `xorm:"index(activity_id) INT(10)"`
	EndTime      int       `xorm:"index(activity_id) INT(10)"`
	CreateUser   string    `xorm:"not null VARCHAR(30)"`
	CreateTime   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Status       int       `xorm:"not null TINYINT(1)"`
	Displayorder int       `xorm:"not null default 0 SMALLINT(6)"`
}
