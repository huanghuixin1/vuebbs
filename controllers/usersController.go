package controllers

import (
	"gopkg.in/baa.v1"
	"../db"
	"../dto/returnStatus"
	"../entity"
)

type usersController struct{}

var UsersController = usersController{}

/**
 通过板块id获取列表
 */
func (this *usersController) GetCount(b *baa.Context) {
	userCount := db.UserInfoDAL.GetCount()

	b.JSON(200, returnStatus.ReturnStatus{returnStatus.Ok, userCount})
}


/**
 注册用户信息
 */
func (this *usersController) Regist(b *baa.Context) {
	user := &entity.UserInfo{}
	user.LoginName = b.Req.PostFormValue("LoginName")
	user.LoginPwd = b.Req.PostFormValue("LoginPwd")
	user.EmailAddr = b.Req.PostFormValue("EmailAddr")
	user.NickName = b.Req.PostFormValue("NickName")


	//判断参数
	if user.CheckRegistFiled() {
		b.JSON(200, returnStatus.ReturnStatus{Status:returnStatus.ParamError, Data:nil })
	} else if db.UserInfoDAL.EmailOrLoginNameExist(*user) {
		//查看用户的登录名或者邮箱是否存在
		b.JSON(200, returnStatus.ReturnStatus{Status:returnStatus.UserExist, Data:nil })
	} else {
		//执行添加操作
		db.UserInfoDAL.Regist(user)
		if user.Id <= 0 {
			b.JSON(200, returnStatus.ReturnStatus{Status:returnStatus.Error, Data:nil })    //系统错误
		}

		//添加到session

		//返回用户信息
		b.JSON(200, returnStatus.ReturnStatus{Status:returnStatus.Ok, Data:user})
	}
}