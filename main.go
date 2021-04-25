package main

import (
	"gf-in-action/boot"
	"gf-in-action/library/constant"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcmd"
	"log"

	_ "gf-in-action/app"
)

func main() {
	if boot.RunMode == constant.RunModeCLI {
		if err := gcmd.AutoRun(); err != nil {
			log.Fatal("failed to run cli script, error: " + err.Error())
		}
		return
	}

	g.Server().Run()
}
