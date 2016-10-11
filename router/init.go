package router

import (
	"gopkg.in/baa.v1"
	"github.com/go-baa/cache"
	_ "github.com/go-baa/cache/redis"
	//"github.com/go-baa/render"
	"github.com/go-baa/render"
	"os"
	"path"
	"io/ioutil"
	"strings"
	"fmt"
)

func Init(b *baa.Baa) {
	b.Use(crossHeader())

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

	b.SetDI("config",initConfig())
}

func crossHeader() baa.HandlerFunc {
	return func(c *baa.Context) {
		c.Resp.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

//初始化配置文件
func initConfig() map[string]string{
	currPath, _ := os.Getwd()
	file, _ := os.Open(path.Join(currPath, "conf/app.ini"))
	defer file.Close()
	resu, _ := ioutil.ReadAll(file)
	strResu := string(resu);

	splits := strings.Split(strResu, "\n")

	ret := make(map[string]string)
	for _, v := range splits {

		if strings.Trim(v, " ") == "" || v[0] == '#'{
			continue
		}
		keyval := strings.Split(v, "=")

		ret[strings.Trim(keyval[0], " ")] = strings.Trim(strings.Trim(keyval[1], " "), "\"")
	}

	fmt.Println(ret)
	return ret
}
