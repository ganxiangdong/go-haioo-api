package model

import (
	"time"
)

type HaiouSnsTags struct {
	Id            int       `xorm:"not null pk autoincr INT(11)"`
	Tagname       string    `xorm:"not null index(tagname) VARCHAR(100)"`
	Brandid       int       `xorm:"INT(10)"`
	Brand         string    `xorm:"default '' VARCHAR(80)"`
	Keywords      string    `xorm:"VARCHAR(500)"`
	Description   string    `xorm:"default '' VARCHAR(500)"`
	FavoriteCount int       `xorm:"not null default 0 INT(11)"`
	ShareCount    int       `xorm:"not null default 0 INT(11)"`
	Pic           string    `xorm:"VARCHAR(500)"`
	CreateTime    time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	CreateUser    string    `xorm:"VARCHAR(30)"`
	Status        int       `xorm:"not null default 1 index(tagname) TINYINT(1)"`
}
