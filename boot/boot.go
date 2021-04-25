package boot

import (
	"gf-in-action/library/constant"
	_ "gf-in-action/packed"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcmd"
	"github.com/gogf/gf/os/gfile"
	"os"
)

// 系统运行时变量
var (
	ENV      string // 系统运行环境. dev: 开发环境, test: 测试环境, prod: 生产环境
	RunMode  string // 系统运行模式. WEB: web 方式运行, CLI: 命令行模式运行
	RootPath string // 系统运行目录.
)

func init() {
	// 设置运行模式
	if gcmd.ContainsOpt("c") || gcmd.ContainsOpt("cli") {
		RunMode = constant.RunModeCLI
	} else {
		RunMode = constant.RunModeWeb
	}

	// 设置项目环境
	_ = os.Setenv("GF_IN_ACTION_ENV", gcmd.GetOpt("env", constant.ENVDev))
	ENV = os.Getenv("GF_IN_ACTION_ENV")

	// 设置配置文件
	g.Cfg().SetFileName("config_" + ENV + ".toml")

	// 获取运行目录
	RootPath = gfile.SelfDir()
}
