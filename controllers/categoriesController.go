package controllers

import (
	"gopkg.in/baa.v1"
	"../db"
	"../dto/returnStatus"
)

type categoriesController struct{}

var CategoriesController = categoriesController{}


//根据板块获取帖子数量
func (this *categoriesController) GetNameByCid(b *baa.Context) {
	cid := b.QueryInt("cid")

	cname := db.CategoriesDAL.GetNameById(cid)

	b.JSON(200, returnStatus.ReturnStatus{Status:returnStatus.Ok, Data:cname})
}

//获取板块列表名字和主题数量
func (this *categoriesController) GetListName(b *baa.Context) {
	resu := db.CategoriesDAL.GetListName()

	b.JSON(200, returnStatus.ReturnStatus{Status:returnStatus.Ok, Data:resu})
}
