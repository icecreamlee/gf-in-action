package main

import (
	_ "gf-in-action/boot"
	_ "gf-in-action/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
