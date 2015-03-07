package model

import (
	"time"
)

type HaiouProductCat struct {
	Catid          int       `xorm:"not null pk autoincr INT(9)"`
	Cat            string    `xorm:"VARCHAR(50)"`
	Title          string    `xorm:"TEXT"`
	Keyword        string    `xorm:"TEXT"`
	Description    string    `xorm:"TEXT"`
	Isindex        int       `xorm:"default 0 TINYINT(1)"`
	CharIndex      string    `xorm:"CHAR(1)"`
	AllChar        string    `xorm:"VARCHAR(50)"`
	Pic            string    `xorm:"VARCHAR(150)"`
	Brand          string    `xorm:"VARCHAR(500)"`
	ExtTable       string    `xorm:"VARCHAR(30)"`
	ExtFieldCat    int       `xorm:"INT(11)"`
	TaxRate        float64   `xorm:"DOUBLE(8,3)"`
	Hscode         string    `xorm:"default '' VARCHAR(30)"`
	WaitingTimeEnd string    `xorm:"default '' VARCHAR(30)"`
	WaitingTime    string    `xorm:"default '' VARCHAR(30)"`
	Displayorder   int       `xorm:"not null SMALLINT(6)"`
	CreateUser     string    `xorm:"VARCHAR(30)"`
	CreateTime     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
