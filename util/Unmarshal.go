package util

import (
	"BitcoinProject/entity"
	"encoding/json"
	"fmt"
)

func GetResult(body []byte)*entity.RPCResult  {
	RPCResult:=new(entity.RPCResult)
	err :=json.Unmarshal(body,RPCResult)
	if err!=nil{
		fmt.Println(err.Error())
		return nil
	}
	return RPCResult

}
