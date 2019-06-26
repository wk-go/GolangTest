package routers

import (
	"github.com/astaxie/beego"
	"golang_test/framework/beego/WebIM/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
