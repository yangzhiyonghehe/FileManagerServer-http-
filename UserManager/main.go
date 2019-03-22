package main

import (
	_ "FileManagerServer/UserManager/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

