package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	c.TplName = "index.tpl"
}



type RegisterController struct{
	beego.Controller
}

func (c *RegisterController) Get(){
	c.TplName = "register.tpl"
}