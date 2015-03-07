package model

type HaiouMemberCart struct {
	Id            int     `xorm:"not null pk autoincr INT(10)"`
	MemberId      int     `xorm:"not null default 0 index INT(10)"`
	Phoneid       string  `xorm:"not null default '0' index VARCHAR(30)"`
	ProductNumber int     `xorm:"not null TINYINT(3)"`
	Price         string  `xorm:"not null DECIMAL(10,2)"`
	Weight        float64 `xorm:"not null DOUBLE(10,3)"`
	UpdateTime    int     `xorm:"not null INT(10)"`
}
