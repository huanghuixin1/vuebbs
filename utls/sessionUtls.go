package utls

import (
	//"github.com/go-baa/cache"
	"gopkg.in/baa.v1"
	//"github.com/satori/go.uuid"
	"github.com/satori/go.uuid"
	"time"
	"net/http"
	"github.com/go-baa/cache"
	"fmt"
	"strings"
)

type sessionUtls struct {
}

var SessionUtls = sessionUtls{}

const sessionExpirSecond = 60 * 60 * 2
const sessionId = "vuebbs_sessionid"

//从cookie中获取或创建 sessionid的值
func (this sessionUtls) getSessionIdCookie(b *baa.Context) *http.Cookie {
	var sessionVal string
	//从cookie集合中拿到sessionid的cookie
	cookie, errCookie := b.Req.Cookie(sessionId)
	if errCookie != nil {
		//如果没有获取到  则 创建一个新的
		sessionVal = strings.Replace(uuid.NewV4().String(), "-", "", -1)
		//重新创建一个cookie
		cookie = &http.Cookie{
			Name: sessionId,
			Value:sessionVal,
		}
	}

	b.Req.AddCookie(cookie)

	return cookie
}

//重置session过期时间
func (this sessionUtls) RefreshSessionExpir(b *baa.Context) {
	//拿到当前用户的sessionid对应的cookie
	cookie := this.getSessionIdCookie(b)
	//重置cookie的过期时间
	cookie.Expires = time.Now().Add(sessionExpirSecond * time.Second)
	b.Resp.Header().Add("set-cookie", cookie.String())

	//从缓存中拿到对应的值
	cacher := b.DI("cache").(cache.Cacher)
	sessionKeyVal := this.getSessionKeyVal(b)
	cacher.Get(cookie.Value, &sessionKeyVal)

	//然后给缓存重置过期时间
	cacher.Set(cookie.Value, sessionKeyVal, sessionExpirSecond)
}

//获取当前session的所有键值对(map)
func (this sessionUtls) getSessionKeyVal(b *baa.Context) map[string]interface{} {
	sessionIdVal := this.getSessionIdCookie(b).Value
	cacher := b.DI("cache").(cache.Cacher)
	var sessionKeyVal map[string]interface{}
	cacher.Get(sessionIdVal, &sessionKeyVal)
	if sessionKeyVal == nil {
		sessionKeyVal = make(map[string]interface{})
	}
	return sessionKeyVal
}

//设置值
func (this sessionUtls) SetData(key string, val interface{}, b *baa.Context) error {
	//从缓存中拿到对应的值
	cacher := b.DI("cache").(cache.Cacher)
	sessionKeyVal := this.getSessionKeyVal(b)

	//重置设置值
	sessionKeyVal[key] = val
	//拿到sessionId的值
	sessionIdVal := this.getSessionIdCookie(b).Value
	err := cacher.Set(sessionIdVal, sessionKeyVal, sessionExpirSecond)
	if err != nil{
		fmt.Println(err)
	}
	return err
}

//获取值
func (this sessionUtls) GetData(key string, b *baa.Context) interface{} {
	kv := this.getSessionKeyVal(b)
	if kv == nil {
		return nil
	} else {
		return kv[key]
	}
}