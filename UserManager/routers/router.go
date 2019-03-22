package routers

import (
	"FileManagerServer/UserManager/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/Register", &controllers.RegisterController{})
}
