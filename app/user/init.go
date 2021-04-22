package user

import (
	"gf-in-action/app/common"
	"gf-in-action/app/user/api"
	"gf-in-action/app/user/cli"
	"gf-in-action/app/user/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcmd"
)

func init() {
	if common.RunMode == common.RunModeCLI {
		setCMDRouter()
		return
	}
	setWebRouter()
}

func setWebRouter() {
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

func setCMDRouter() {
	err := gcmd.BindHandle("test", cli.User.Test)
	if err != nil {
		panic("failed to bind cmd handle, error:" + err.Error())
	}
}
