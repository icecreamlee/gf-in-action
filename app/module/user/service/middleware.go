package service

import (
	"gf-in-action/app/module/user/define"
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
		response.Json(r, define.ErrCodeLoginFirst, define.ErrMsgLoginFirst)
		return
	}

	// 设置上下文参数，如用户ID、用户名等其他需要在请求整个生命周期用到的参数
	Context.UserId = int(user.Id)
	Context.Username = user.Username
	r.SetCtxVar(ContextKey, Context)

	r.Middleware.Next()
}
