package routers

import (
	"git.oschina.net/walkskyer/GolangTest/framework/beego/proj1/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
