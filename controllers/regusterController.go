package controllers

import (
	"BitcoinProject/models"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Get() {
	//设置login.html为模板文件
	r.TplName = "registers.html"
}
func (r *RegisterController) Post() {
	//	1.解析请求数据
	var user models.User
	err := r.ParseForm(&user)
	if err != nil {
		//返回错误信息给浏览器，提示用户
		r.Ctx.WriteString("抱歉，解析数据错误")
		return
	}
	//	2.保存用户信息到数据库
	_, err = user.SaverUser()
	//	3.返回前端结果（成功跳登录页面，失败弹出错误信息）
	if err != nil {
		//fmt.Println(err.Error())
		r.Ctx.WriteString("抱歉，用户注册失败，请重试！")
		return
	}
	//	 用户注册成功
	r.TplName = "index.html"
}
