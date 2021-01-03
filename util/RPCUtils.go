package util

import (
	"BitcoinProject/entity"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//准备一个json数据

func PrepareJson(method string, params interface{}) []byte {
	rpcReq := entity.RPCRequest{
		Id:      time.Now().Unix(),
		Method:  method,
		Jsonrpc: RPCVERSION,
		Params:  params,
	}
	reqBytes, err := json.Marshal(&rpcReq)
	if err != nil {
		fmt.Println(err.Error())
		return reqBytes
	}
	//fmt.Println("准备好的json数据：", string(reqBytes))
	return reqBytes
}
func Dopost(url string, reqBytes []byte) []byte {
	client := &http.Client{}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBytes))
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	for key, value := range ReqHeader() {
		request.Header.Add(key,value)
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	body, err := ioutil.ReadAll(response.Body)
	code := response.StatusCode
	if code == 200 {
		fmt.Println("请求成功")
	} else {
		fmt.Println(code)
		fmt.Println("请求失败")
	}
	return body
}
