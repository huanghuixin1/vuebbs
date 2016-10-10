package controllers

import (
"gopkg.in/baa.v1"
"../db"
"../dto/returnStatus"
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