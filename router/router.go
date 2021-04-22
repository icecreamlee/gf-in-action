package router

import (
	_ "gf-in-action/app/user"
	"gf-in-action/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.BindStatusHandler(404, func(r *ghttp.Request) {
		_ = r.Response.WriteJson(response.JsonRes{
			Code:    404,
			Message: "Page Not Found.",
			Data:    g.Map{},
		})
	})
}
