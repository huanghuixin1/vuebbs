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

//检测注册信息中的字段
func (this UserInfo) CheckRegistFiled() bool {

	if this.NickName == "" ||
		this.LoginName == "" ||
		this.LoginPwd == "" ||
		this.EmailAddr == ""{
		return true
	}else{
		return false
	}
}



func (UserInfo) TableName() string {
	return "userinfos"
}