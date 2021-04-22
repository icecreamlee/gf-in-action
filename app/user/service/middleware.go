package service

import (
	"gf-in-action/library/response"
	"github.com/gogf/gf/net/ghttp"
)

// Middleware 用户模块中间件变量
var Middleware = middleware{}

// middlewareService 用户模块中间件
type middleware struct{}

// Auth 登录校验中间件
func (m middleware) Auth(r *ghttp.Request) {
	// 校验登录
	uid := r.Session.GetInt("uid")
	user, err := User.UserInfo(uid)
	if err != nil || user == nil {
		response.Json(r, 1, ErrLoginFirst.Error())
		return
	}

	// 设置上下文参数
	Context.UserId = int(user.Id)
	Context.Username = user.Username
	r.SetCtxVar(ContextKey, Context)

	r.Middleware.Next()
}
