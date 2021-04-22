package main

import (
	"gf-in-action/app/common"
	_ "gf-in-action/boot"
	_ "gf-in-action/router"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcmd"
	"log"
)

func main() {
	if common.RunMode == common.RunModeCLI {
		if err := gcmd.AutoRun(); err != nil {
			log.Fatal("failed to run cli script, error: " + err.Error())
		}
		return
	}

	g.Server().Run()
}
