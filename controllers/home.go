package controllers

import (
	"BitcoinProject/models"
	"BitcoinProject/util"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type Home struct {
	beego.Controller
}

func (h *Home) Post() {
	fmt.Println("是否为ajax请求",h.IsAjax())
	var command models.Command
	data:=h.Ctx.Input.RequestBody
	json.Unmarshal(data,&command)
	fmt.Println(command)
	method := command.Method
	params := command.Params
	commands:= make(map[string]interface{})
	height,_:=strconv.Atoi(params)
	commands["getblockhash"]=util.GetBlockHash(height)
	commands["getmininginfo"]=util.GetMininginfo()
	for key, value := range commands {
		if method==key {
			result,err:=json.Marshal(value)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			h.Data["json"]=string(result)
			h.ServeJSON()
		}
	}

}
