package model

type HaiouNewsData struct {
	Nid int    `xorm:"not null pk INT(11)"`
	Con string `xorm:"not null MEDIUMTEXT"`
}
