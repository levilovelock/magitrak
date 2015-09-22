package main

import (
	_ "github.com/astaxie/beego/session/mysql"
	_ "github.com/levilovelock/magitrak/docs"
	_ "github.com/levilovelock/magitrak/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
	}
	beego.Run()
}
