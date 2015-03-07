package model

type HaiouSearchWord struct {
	Id        int    `xorm:"not null pk autoincr INT(11)"`
	Keyword   string `xorm:"VARCHAR(80)"`
	CharIndex string `xorm:"VARCHAR(80)"`
	Nums      int    `xorm:"default 0 INT(11)"`
}
