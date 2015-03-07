package model

import (
	"time"
)

type HaiouMemberGrade struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	Name       string    `xorm:"not null VARCHAR(50)"`
	Demand     int       `xorm:"default 0 INT(10)"`
	Treatment  string    `xorm:"TEXT"`
	Blogo      string    `xorm:"VARCHAR(255)"`
	Logo       string    `xorm:"VARCHAR(255)"`
	Valid      int       `xorm:"default 0 INT(1)"`
	Sum        int       `xorm:"default 0 INT(11)"`
	Rate       float32   `xorm:"default 0.0 FLOAT(3,1)"`
	CreateUser string    `xorm:"VARCHAR(30)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
