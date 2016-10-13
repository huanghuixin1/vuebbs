package db

import (
	"../entity"
	"database/sql"
	"../utls"
	"../entity/userlevel"
)

type userInfoDAL struct {

}

var UserInfoDAL = userInfoDAL{}


//注册用户
func (this *userInfoDAL)Regist(user *entity.UserInfo) {
	//初始化用户信息
	user.CreateTime = utls.TimeUtls.GetGTMTimeTamp()
	user.Level = userlevel.Normal
	user.IsDelete = false

	db := openDb()
	defer db.Close()

	db.Create(user)
}

//登录方法
func (this *userInfoDAL)Login(loginName string, email string, pwd string) (*entity.UserInfo, error) {
	sqlStr := `SELECT
		id,
		NickName,
		LoginName,
		LoginPwd,
		EmailAddr,
		IsDelete,
		ULevel,
		CreateTime
	FROM userinfos
	WHERE (LoginName = ? OR EmailAddr = ?) AND LoginPwd = ? AND IsDelete = 0`
	db := openDb()
	defer db.Close()
	row := db.Raw(sqlStr, loginName, email, pwd).Row()

	userinfo, err := row2Entity(*row)

	return userinfo, err
}

//根据id获取用户信息
func (this *userInfoDAL) GetById(id int) (*entity.UserInfo) {

	db := openDb()
	defer db.Close()

	user := &entity.UserInfo{}
	db.Where("id = ? and IsDelete = 0", id).First(user)

	return user
}

//获取会员总数
func (this userInfoDAL) GetCount() int {
	db := openDb()
	defer db.Close()

	var count int
	db.Table("userinfos").Where("IsDelete = 0").Count(&count)

	return count
}

func row2Entity(row sql.Row) (*entity.UserInfo, error) {
	userinfo := new(entity.UserInfo)

	err := row.Scan(&userinfo.Id, &userinfo.NickName, &userinfo.LoginName, &userinfo.LoginPwd, &userinfo.EmailAddr,
		&userinfo.IsDelete, &userinfo.Level, &userinfo.CreateTime)

	return userinfo, err
}