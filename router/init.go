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
	"../utls"
	"encoding/gob"
)

func Init(b *baa.Baa) {
	gob.Register(map[string]interface{}{})
	b.Use(crossHeader())
	config := initConfig()
	fmt.Println(config)
	b.SetDI("config", config)
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
		Prefix:   "vuebbs_",
		Adapter:  "redis",
		Config:    map[string]interface{}{
			"host":     config["redishost"],
			"port":     config["redisport"],
			"password": config["redispwd"],
			"poolsize": 10,
		},
	}))
}

//处理跨域和sessionid
func crossHeader() baa.HandlerFunc {
	return func(c *baa.Context) {
		c.Resp.Header().Add("Access-Control-Allow-Origin", "*")
		//刷新session的过期时间
		utls.SessionUtls.RefreshSessionExpir(c)
		c.Next()

	}
}

//初始化配置文件
func initConfig() map[string]string {
	currPath, _ := os.Getwd()
	file, _ := os.Open(path.Join(currPath, "conf/app.ini"))
	defer file.Close()
	resu, _ := ioutil.ReadAll(file)
	strResu := strings.Replace(string(resu), "\r", "", -1);

	splits := strings.Split(strResu, "\n")

	ret := make(map[string]string)
	for _, v := range splits {

		if strings.Trim(v, " ") == "" || v[0] == '#' {
			continue
		}
		keyval := strings.Split(v, "=")

		if len(keyval) >= 2 {
			ret[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], " ")
		}
	}

	fmt.Println(ret)
	return ret
}
