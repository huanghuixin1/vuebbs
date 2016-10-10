package router

import (
	"gopkg.in/baa.v1"
	"github.com/go-baa/cache"
	_ "github.com/go-baa/cache/redis"
	//"github.com/go-baa/render"
	"github.com/go-baa/render"
)

func Init(b *baa.Baa) {
	b.Static(
		"/",
		"html/dist",
		true, nil,
	)


	b.SetDI("render", render.New(render.Options{
		Baa:        b,
		Root:       "html/dist",
		Extensions: []string{".html", ".tmpl"},
	}))

	//配置缓存
	b.SetDI("cache", cache.New(cache.Options{
		Name:     "redis",
		Prefix:   "note_online_",
		Adapter:  "redis",
		Config:    map[string]interface{}{
			"host":     "127.0.0.1",
			"port":     "6379",
			"password": "redis8114359",
			"poolsize": 10,
		},
	}))
}
