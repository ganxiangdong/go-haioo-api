package model

import (
	"time"
)

type HaiouMember struct {
	Userid         int       `xorm:"not null pk autoincr INT(8)"`
	Phoneid        string    `xorm:"VARCHAR(30)"`
	Username       string    `xorm:"index(user) CHAR(30)"`
	Password       string    `xorm:"index(user) CHAR(32)"`
	OpenType       int       `xorm:"INT(3)"`
	OpenSecretid   string    `xorm:"VARCHAR(80)"`
	Name           string    `xorm:"VARCHAR(30)"`
	Sex            int       `xorm:"TINYINT(1)"`
	Mobile         string    `xorm:"index VARCHAR(18)"`
	Email          string    `xorm:"index VARCHAR(50)"`
	Provinceid     int       `xorm:"INT(11)"`
	Cityid         int       `xorm:"INT(11)"`
	Areaid         int       `xorm:"INT(11)"`
	Area           string    `xorm:"VARCHAR(255)"`
	Logo           string    `xorm:"VARCHAR(520)"`
	Ip             string    `xorm:"CHAR(15)"`
	Points         int       `xorm:"INT(10)"`
	Status         int       `xorm:"index TINYINT(1)"`
	Grade          int       `xorm:"default 1 INT(3)"`
	UpdateDate     int       `xorm:"default 0 INT(10)"`
	Regtime        time.Time `xorm:"DATETIME"`
	Chkstatus      int       `xorm:"default 0 TINYINT(1)"`
	ClientPlatform int       `xorm:"TINYINT(1)"`
	ClientSystem   int       `xorm:"TINYINT(1)"`
	Lastlogintime  int       `xorm:"INT(10)"`
	CreateTime     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
