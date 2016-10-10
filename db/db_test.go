package db

import (
	"testing"
	"time"
)

//[UserInfoDAL.Login]
func Test_userLogin(t *testing.T) {
	user, err := UserInfoDAL.Login("hhx", "", "123")
	if err != nil {
		t.Error(err)
	} else {
		t.Log("查询完成", user)
	}
}

//[ArticlesDAL.GetListByCategoryId]
func Test_GetArticleListByCid(t *testing.T) {
	ret, err := ArticlesDAL.GetListByCategoryId(0, 0, 2, 0)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(len(ret), ret)
	}
}

//[UserInfoDAL.GetById]测试获取用户信息
func Test_getuser(t *testing.T) {
	t.Log(UserInfoDAL.GetById(1))
	t.Log(UserInfoDAL.GetById(1111))
}

//时间戳测试
func Test_timetamp(t *testing.T) {
	now := time.Now()
	_, offset := now.Zone()
	timetamp := time.Now().Unix() - int64(offset)

	t.Log(timetamp)
}