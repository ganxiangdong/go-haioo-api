package model

import (
	"time"
)

type HaiouAdvs struct {
	Id         int       `xorm:"not null pk autoincr INT(4)"`
	Width      string    `xorm:"VARCHAR(10)"`
	Height     string    `xorm:"VARCHAR(10)"`
	AdType     int       `xorm:"not null default 1 TINYINT(1)"`
	Name       string    `xorm:"not null default '' VARCHAR(50)"`
	Group      string    `xorm:"VARCHAR(50)"`
	Con        string    `xorm:"MEDIUMTEXT"`
	Date       time.Time `xorm:"DATETIME"`
	Price      string    `xorm:"DECIMAL(10,2)"`
	Unit       string    `xorm:"not null ENUM('day','week','month')"`
	Total      int       `xorm:"default 0 TINYINT(4)"`
	CreateUser string    `xorm:"VARCHAR(30)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
