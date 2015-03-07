package model

type HaiouPaymentType struct {
	PaymentId         int    `xorm:"not null pk autoincr TINYINT(3)"`
	PaymentType       string `xorm:"VARCHAR(20)"`
	PaymentName       string `xorm:"VARCHAR(100)"`
	PaymentCommission string `xorm:"default '0' VARCHAR(8)"`
	PaymentLogo       string `xorm:"VARCHAR(255)"`
	PaymentDesc       string `xorm:"TEXT"`
	PaymentConfig     string `xorm:"TEXT"`
	Active            int    `xorm:"default 0 TINYINT(1)"`
	Nums              int    `xorm:"default 0 TINYINT(3)"`
}
