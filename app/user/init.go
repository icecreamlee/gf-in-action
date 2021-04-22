package user

import (
	"gf-in-action/app/user/api"
	"gf-in-action/app/user/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/user", func(group *ghttp.RouterGroup) {
		group.ALL("/", api.User.Index)
		group.ALL("/index", api.User.Index)
		group.ALL("/login", api.User.Login)
		group.ALL("/register", api.User.Register)
		group.ALL("/logout", api.User.Logout)
	})

	s.Group("/user", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware.Auth)
		group.ALL("/info", api.User.Info)
	})
}
