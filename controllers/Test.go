package controllers

import (
	"BitcoinProject/models"
	"BitcoinProject/util"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type Test struct {
	beego.Controller
}
func (t *Test) Post() {
	//设置login.html为模板文件
	var command models.Command
	err:=t.ParseForm(&command)
	if err != nil {
		fmt.Println(err.Error())
		return

	}
	method:=command.Method
	params:= command.Params
	fmt.Println(params)
	commands:=make(map[string]interface{})
	height,_:=strconv.Atoi(params)
	commands["getblockhash"]=util.GetBlockHash(height)
	for key, value := range commands {
		if key==method {
			t.Data["Result"]=value
			t.Data["Time"]=height
			t.TplName="response_test.html"
		}
		continue
	}
}