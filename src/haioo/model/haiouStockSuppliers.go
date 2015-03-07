package model

type HaiouStockSuppliers struct {
	Id      int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Title   string `xorm:"VARCHAR(40)"`
	Linkman string `xorm:"VARCHAR(20)"`
	Phone   string `xorm:"VARCHAR(20)"`
	Address string `xorm:"VARCHAR(50)"`
	Memo    string `xorm:"TEXT"`
}
