package model

type HaiouMemberPushRelation struct {
	Id            int     `xorm:"not null pk autoincr INT(10)"`
	MemberId      int     `xorm:"not null index(member_id) INT(11)"`
	Phoneid       string  `xorm:"default '' VARCHAR(30)"`
	PushChannelId string  `xorm:"not null default '' index(member_id) VARCHAR(30)"`
	PushUserId    string  `xorm:"not null default '' index(member_id) VARCHAR(72)"`
	ClientSystem  int     `xorm:"default 1 TINYINT(1)"`
	ClientVersion float32 `xorm:"default 0.0 FLOAT(3,1)"`
}
