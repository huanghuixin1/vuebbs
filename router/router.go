package router

import (
	//"git.vodjk.com/golang/common/util"
	"gopkg.in/baa.v1"
	//"github.com/go-baa/cache"
	//"../entity"
	//"fmt"
	//"net/http"

	"../controllers"
	"encoding/json"
)

func Router(app *baa.Baa) {
	app.SetAutoHead(true)
	app.SetAutoTrailingSlash(true)


	//procHtml(app)

	app.Get("/", func(c *baa.Context) {
		c.Redirect(301, "/app.html")
	})

	//app.Get("/userinfo", func(c *baa.Context) {
	//	a := c.Query("a")
	//	c.String(200, a + ",1")
	//})

	//app.Get("/test2", func(c *baa.Context) {
	//	var listUser [2]*entity.UserInfo
	//
	//
	//	c.Set("data", listUser)
	//	//u:= &entity.UserInfo{Age:11, Name:"hhx"}
	//	//c.Set("data", u)
	//
	//	c.HTML(200, "userInfo/index")
	//})

	//app.Get("/test", func(c *baa.Context) {
	//	ca := c.DI("cache").(cache.Cacher)
	//
	//	ca.Set("test", "baa", 100)
	//	c.String(200, "ok")
	//})

	app.Get("/test", func(b * baa.Context) {
		slice := make([]string,0, 100)
		slice = append(slice, "asd")
		resu , _ := json.Marshal(slice)
		b.JSON(200, string(resu))
	})

	app.Group("/articlesApi", func() {
		app.Get("/getbycid", controllers.ArticlesController.GetByCid)
		app.Get("/count",controllers.ArticlesController.GetCoundByCid)
	})

	app.Group("/userinfosApi", func() {
		app.Get("/count", controllers.UsersController.GetCount)
	})

	app.Group("/categoriesApi", func() {
		app.Get("/getNameById", controllers.CategoriesController.GetNameByCid)
	})
}