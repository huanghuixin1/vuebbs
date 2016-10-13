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
func (this *usersController) Regist(b *baa.Context){
	user := &entity.UserInfo{}
	user.LoginName = b.Req.PostFormValue("LoginName")
	user.LoginPwd = b.Req.PostFormValue("LoginPwd")
	user.EmailAddr = b.Req.PostFormValue("EmailAddr")
	user.NickName = b.Req.PostFormValue("NickName")

	//查看用户是否已经存在

}