package model

type HaiouMemberOpinionsType struct {
	Id           int    `xorm:"not null pk autoincr TINYINT(2)"`
	TypeName     string `xorm:"not null default '' VARCHAR(100)"`
	DisplayOrder int    `xorm:"not null TINYINT(2)"`
	Status       int    `xorm:"not null default 1 TINYINT(1)"`
}
