package model

import (
	"github.com/Unknwon/goconfig"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//从库
var XormRead *xorm.Engine

//主库
var XormWriter *xorm.Engine

func init() {
	initXorm()
}

//初使化xorm
func initXorm() {
	var err error
	//获取数据库配置信息
	cfg, err := goconfig.LoadConfigFile("conf/config.ini")
	if err != nil {
		panic("读取config.ini文件失败：")
	}
	readHost, _ := cfg.GetValue("db", "readHost")
	readUser, _ := cfg.GetValue("db", "readUser")
	readPwd, _ := cfg.GetValue("db", "readPwd")
	readDbName, _ := cfg.GetValue("db", "readDbName")
	writerHost, _ := cfg.GetValue("db", "writerHost")
	writerUser, _ := cfg.GetValue("db", "writerUser")
	writerPwd, _ := cfg.GetValue("db", "writerPwd")
	writerDbName, _ := cfg.GetValue("db", "writerDbName")
	//初使化xorm的engine
	XormRead, err = xorm.NewEngine("mysql", readUser+":"+readPwd+"@tcp("+readHost+")/"+readDbName+"?charset=utf8")
	if err != nil {
		panic("连接从库失败")
	}

	XormWriter, err = xorm.NewEngine("mysql", writerUser+":"+writerPwd+"@tcp("+writerHost+")/"+writerDbName+"?charset=utf8")
	if err != nil {
		panic("连接主库失败")
	}

	//配置xorm
	XormRead.ShowSQL = true
	XormWriter.ShowSQL = false
}
