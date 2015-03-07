package model

import (
	"time"
)

type HaiouUserGroup struct {
	GroupId   int       `xorm:"not null pk autoincr INT(11)"`
	Name      string    `xorm:"not null VARCHAR(50)"`
	Des       string    `xorm:"TEXT"`
	Con       string    `xorm:"TEXT"`
	Logo      string    `xorm:"VARCHAR(100)"`
	Minilogo  string    `xorm:"VARCHAR(200)"`
	Statu     int       `xorm:"default 1 TINYINT(4)"`
	CreatTime time.Time `xorm:"DATE"`
	Groupfee  float32   `xorm:"default 0 FLOAT"`
}
