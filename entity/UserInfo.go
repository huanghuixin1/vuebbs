//用户信息类
package entity

type UserInfo struct {
	BaseEntity
	NickName  string `gorm:"column:NickName"`
	LoginName string  `gorm:"column:LoginName"`
	LoginPwd  string `gorm:"column:LoginPwd"`
	EmailAddr string `gorm:"column:EmailAddr"`
	Level     int `gorm:"column:ULevel"`
}

func (UserInfo) TableName() string {
	return "userinfos"
}