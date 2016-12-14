package routers

import (
	"git.oschina.net/walkskyer/GolangTest/framework/beego/WebIM/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
