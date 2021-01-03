package routers

import (
	"BitcoinProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//请求登陆页面接口
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/index.html",&controllers.LoginController{})

	//请求注册页面接口
	beego.Router("/user_register", &controllers.RegisterController{})
	beego.Router("/registers.html", &controllers.RegisterController{})



	beego.Router("/test",&controllers.Test{})


}
