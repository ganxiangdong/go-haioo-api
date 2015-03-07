package model

type HaiouTgOrder struct {
	Id              int    `xorm:"not null pk autoincr INT(11)"`
	OrderId         string `xorm:"not null VARCHAR(15)"`
	MemberId        int    `xorm:"INT(11)"`
	MemberName      string `xorm:"not null VARCHAR(50)"`
	TgId            int    `xorm:"not null INT(11)"`
	TgName          string `xorm:"not null VARCHAR(80)"`
	TgPic           string `xorm:"VARCHAR(80)"`
	Contact         string `xorm:"VARCHAR(30)"`
	Address         string `xorm:"VARCHAR(200)"`
	Tel             string `xorm:"VARCHAR(20)"`
	Remark          string `xorm:"VARCHAR(200)"`
	AdminRemark     string `xorm:"VARCHAR(200)"`
	Price           string `xorm:"default 0.00 DECIMAL(10,2)"`
	Quantity        string `xorm:"not null VARCHAR(10)"`
	CreateTime      int    `xorm:"INT(10)"`
	PaymentTime     int    `xorm:"INT(10)"`
	ShippingTime    int    `xorm:"INT(10)"`
	FinishedTime    int    `xorm:"INT(10)"`
	Status          int    `xorm:"default 20 TINYINT(2)"`
	ShippingName    string `xorm:"VARCHAR(50)"`
	ShippingAddress string `xorm:"VARCHAR(255)"`
	ShippingTel     string `xorm:"VARCHAR(20)"`
	ShippingCompany string `xorm:"VARCHAR(50)"`
	ShippingCode    string `xorm:"VARCHAR(50)"`
}
