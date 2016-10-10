package controllers

import (
	"gopkg.in/baa.v1"
)

type indexController struct{}

var IndexController = indexController{}

func (this *indexController)Index(b *baa.Context) {
	b.HTML(200, "app.html")
}

/*

	app.Get("/userinfo", func(c *baa.Context) {
		a := c.Query("a")
		c.String(200, a + ",1")
	})

 */