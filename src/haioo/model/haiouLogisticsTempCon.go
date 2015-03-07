package model

import (
	"time"
)

type HaiouLogisticsTempCon struct {
	Id           int       `xorm:"not null pk autoincr INT(11)"`
	TempId       int       `xorm:"index INT(11)"`
	DefaultNum   float64   `xorm:"default 0.00 DOUBLE(5,2)"`
	DefaultPrice string    `xorm:"default 0.00 DECIMAL(5,2)"`
	AddNum       float64   `xorm:"default 0.00 DOUBLE(5,2)"`
	AddPrice     string    `xorm:"default 0.00 DECIMAL(5,2)"`
	Free         string    `xorm:"default 0.00 DECIMAL(5,2)"`
	DefineCitys  string    `xorm:"TEXT"`
	CreateUser   string    `xorm:"VARCHAR(30)"`
	CreateTime   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
