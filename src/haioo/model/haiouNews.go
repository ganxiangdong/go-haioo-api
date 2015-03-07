package model

type HaiouNews struct {
	Nid          int    `xorm:"not null pk autoincr INT(11)"`
	Classid      int    `xorm:"not null default 1 SMALLINT(6)"`
	Title        string `xorm:"not null VARCHAR(200)"`
	Ftitle       string `xorm:"not null VARCHAR(120)"`
	Keyboard     string `xorm:"not null VARCHAR(100)"`
	Titleurl     string `xorm:"not null default '0' VARCHAR(200)"`
	Isrec        int    `xorm:"not null default 0 TINYINT(1)"`
	Istop        int    `xorm:"not null default 0 TINYINT(1)"`
	Ispass       int    `xorm:"not null default 0 TINYINT(1)"`
	Firsttitle   int    `xorm:"not null default 0 TINYINT(1)"`
	Onclick      int    `xorm:"not null default 0 INT(11)"`
	Titlefont    string `xorm:"not null VARCHAR(100)"`
	Uid          int    `xorm:"not null INT(11)"`
	Uptime       int    `xorm:"not null INT(12)"`
	Smalltext    string `xorm:"not null VARCHAR(255)"`
	Writer       string `xorm:"not null VARCHAR(50)"`
	Source       string `xorm:"not null VARCHAR(100)"`
	Titlepic     string `xorm:"not null VARCHAR(200)"`
	Ispic        int    `xorm:"not null default 0 TINYINT(1)"`
	Isgid        int    `xorm:"not null default 0 TINYINT(1)"`
	Ispl         int    `xorm:"not null TINYINT(1)"`
	Userfen      int    `xorm:"not null SMALLINT(6)"`
	Newstempid   string `xorm:"not null VARCHAR(40)"`
	Pagenum      int    `xorm:"not null default 0 INT(11)"`
	Lastedittime int    `xorm:"not null INT(12)"`
	Vote         string `xorm:"not null default '0' CHAR(255)"`
	Special      string `xorm:"not null default '0' CHAR(255)"`
	Products     string `xorm:"not null CHAR(255)"`
	ImgsUrl      string `xorm:"not null TEXT"`
	VideosUrl    string `xorm:"not null TEXT"`
	Admin        string `xorm:"VARCHAR(50)"`
}
