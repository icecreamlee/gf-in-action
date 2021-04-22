package boot

import (
	"gf-in-action/app/common"
	_ "gf-in-action/packed"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcmd"
	"github.com/gogf/gf/os/gfile"
	"os"
)

func init() {
	// 设置运行模式
	if gcmd.ContainsOpt("c") || gcmd.ContainsOpt("cli") {
		common.RunMode = common.RunModeCLI
	} else {
		common.RunMode = common.RunModeWeb
	}

	// 设置项目环境
	if gcmd.ContainsOpt("env") {
		_ = os.Setenv("SMS_GO_ENV", gcmd.GetOpt("env", common.ENVProd))
	}
	common.ENV = os.Getenv("SMS_GO_ENV")

	// 设置配置文件
	g.Cfg().SetFileName("config_" + common.ENV + ".toml")

	// 获取运行目录
	common.RootPath = gfile.SelfDir()
}
