package app

import (
	_ "gf-in-action/app/module/user"
	"gf-in-action/boot"
	"gf-in-action/library/constant"
	"gf-in-action/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	if boot.RunMode == constant.RunModeWeb {
		s := g.Server()
		s.BindStatusHandler(404, func(r *ghttp.Request) {
			_ = r.Response.WriteJson(response.JsonRes{
				Code:    404,
				Message: "Page Not Found.",
				Data:    g.Map{},
			})
		})
	}
}
