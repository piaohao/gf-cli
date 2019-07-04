package main

import (
	_ "github.com/gogf/gf-cli/boot"
	_ "github.com/gogf/gf-cli/router"
	"github.com/gogf/gf/g"
)

func main() {
	g.Server().Run()
}
