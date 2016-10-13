package controllers

import (
	"gopkg.in/baa.v1"
	"../db"
	"../dto/returnStatus"
)

type articlesController struct{}

var ArticlesController = articlesController{}

/**
 通过板块id获取列表
 */
func (this *articlesController) GetByCid(b *baa.Context) {
	cid := b.QueryInt("cid") //板块id
	size := b.QueryInt("size")  //最大数量
	minId := b.QueryInt("minId") // 取大于该id的数据
	maxId := b.QueryInt("maxId") //  取小于该id的数据

	if size <= 0 {
		size = 10
	}

	ret, err := db.ArticlesDAL.GetListByCategoryId(cid, size, minId, maxId)
	if err != nil {
		b.JSON(200, returnStatus.ReturnStatus{Status:returnStatus.Error, Data:err })
	}
	returnData := returnStatus.ReturnStatus{
		Status:returnStatus.Ok,
		Data:ret,
	}

	b.JSON(200, returnData)
}

//根据板块获取帖子数量
func (this *articlesController) GetCoundByCid(b *baa.Context) {
	cid := b.QueryInt("cid")

	count := db.ArticlesDAL.GetCountByCid(cid)

	b.JSON(200, returnStatus.ReturnStatus{Status:returnStatus.Ok, Data:count})
}

//获取详情
func (this *articlesController) GetDetail(b *baa.Context) {
	id := b.QueryInt("id")

	ret, err := db.ArticlesDAL.GetDetail(id)

	var returnRet returnStatus.ReturnStatus

	if err == nil {
		returnRet = returnStatus.ReturnStatus{Status:returnStatus.Ok, Data:ret}
	} else {
		returnRet = returnStatus.ReturnStatus{Status:returnStatus.Error, Data:err}
	}

	b.JSON(200, returnRet)
}