package router

import (
	//"git.vodjk.com/golang/common/util"
	"gopkg.in/baa.v1"
	//"github.com/go-baa/cache"
	//"../entity"
	//"fmt"
	//"net/http"
	"github.com/baa-middleware/gzip"
	"../controllers"
	"../utls/sessionUtls"
	"time"
	"fmt"
)

func Router(app *baa.Baa) {
	app.SetAutoHead(true)
	app.SetAutoTrailingSlash(true)

	//压缩比例
	app.Use(gzip.Gzip(gzip.Options{CompressionLevel: 6}))

	//转到app.html
	app.Get("/", func(c *baa.Context) {
		c.Redirect(301, "/app.html")
	})

	//测试使用
	app.Get("/test", func(b *baa.Context) {
		fmt.Println(sessionUtls.SessionUtls.GetData("aaa",b))
		sessionUtls.SessionUtls.SetData("aaa", time.Now(), b)
		b.String(200, "test")
	})

	//帖子信息相关
	app.Group("/articlesApi", func() {
		app.Get("/getbycid", controllers.ArticlesController.GetByCid) //通过板块id获取列表
		app.Get("/count", controllers.ArticlesController.GetCoundByCid) // 获取对应板块的帖子爽
		app.Get("/getDetail", controllers.ArticlesController.GetDetail) //获取详情
	})

	//用户相关
	app.Group("/userinfosApi", func() {
		app.Get("/count", controllers.UsersController.GetCount)// 获取用户总数
		app.Get("/getCurrent", controllers.UsersController.GetCurrent)// 获取当前用户信息
		app.Post("/regist", controllers.UsersController.Regist)//注册用户
	})

	//板块相关
	app.Group("/categoriesApi", func() {
		app.Get("/getNameById", controllers.CategoriesController.GetNameByCid)// 根据id获取板块名字
		app.Get("/getListName", controllers.CategoriesController.GetListName) //获取板块列表
	})
}