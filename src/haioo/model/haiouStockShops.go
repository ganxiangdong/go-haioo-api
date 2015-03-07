package model

type HaiouStockShops struct {
	Id    int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Title string `xorm:"VARCHAR(40)"`
}
