package user

import (
	"gf-in-action/app/module/user/api"
	"gf-in-action/app/module/user/cli"
	"gf-in-action/app/module/user/cron"
	"gf-in-action/app/module/user/service"
	"gf-in-action/boot"
	"gf-in-action/library/constant"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcmd"
	"github.com/gogf/gf/os/gcron"
)

// init 初始化当前系统模块中需要的运行参数，如设置路由
func init() {
	if boot.RunMode == constant.RunModeCLI {
		setCMDRouter()
		return
	}

	setWebRouter()
	setCrontab()
}

// setWebRouter 设置API接口路由
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

// setCrontab 设置定时任务
func setCrontab() {
	var err error
	// 定义crontab，子数组第一个参数为任务名称，第二个为crontab表达式，第三个为任务执行方法，第四个为是否单实例执行(非必填)
	var crontab = [][]interface{}{
		{"UserTest", "0 * * * * *", cron.User.Test, "singleton"},
	}

	for _, c := range crontab {
		if len(c) >= 4 && c[3] == "singleton" {
			_, err = gcron.AddSingleton(c[1].(string), c[2].(func()), c[0].(string))
		} else {
			_, err = gcron.Add(c[1].(string), c[2].(func()), c[0].(string))
		}
		if err != nil {
			panic("failed to add cron, error:" + err.Error())
		}
	}

}

// setCMDRouter 设置命令行程序路由
func setCMDRouter() {
	var err error
	// 定义命令行脚本方法，子数组第一个参数为命令行参数，第二个为脚本执行方法
	var clis = [][]interface{}{
		{"user/test", cli.User.Test},
	}

	for _, c := range clis {
		err = gcmd.BindHandle(c[0].(string), c[1].(func()))
		if err != nil {
			panic("failed to bind cmd handle, error:" + err.Error())
		}
	}
}
