package cacheUtls

import (
	"gopkg.in/baa.v1"
	"github.com/go-baa/cache"
	"strconv"
	"../../entity"
	"../../db"
)

type cacheUtls struct {

}

var CacheUtls = cacheUtls{}

const userInfoKeyPre = "cacheHelper_userInfoKeyPre"

//通过id获取用户信息
func (this cacheUtls) GetUserInfo(id int) *entity.UserInfo {
	finalKey := userInfoKeyPre + strconv.Itoa(id)
	b := baa.Default()
	cacher := b.GetDI("cache").(cache.Cacher)

	user := &entity.UserInfo{}
	cacher.Get(finalKey, user)

	//如果缓存里面没有  那去数据库里找
	if user == nil{
		user = db.UserInfoDAL.GetById(id)
		if user != nil{
			this.SetUserInfo(*user)
		}
		return user
	}else{
		return user
	}
}

//设置用户信息
func (this cacheUtls) SetUserInfo(user entity.UserInfo) {
	finalKey := userInfoKeyPre + strconv.Itoa(user.Id)
	b := baa.Default()
	cacher := b.GetDI("cache").(cache.Cacher)

	cacher.Set(finalKey, user, 9999999)
}