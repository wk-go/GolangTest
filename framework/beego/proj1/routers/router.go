package routers

import (
	"github.com/astaxie/beego"
	"golang_test/framework/beego/proj1/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
