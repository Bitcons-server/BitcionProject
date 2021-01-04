package controllers

import (
	"BitcoinProject/models"
	"BitcoinProject/util"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type LoginController struct {
	beego.Controller
}

//直接访问login.html页面的请求
func (l *LoginController) Get() {
	//设置login.html为模板文件
	l.TplName = "index.html"
}

//用户登录接口
func (l *LoginController) Post() {
		bodyData, err := ioutil.ReadAll(l.Ctx.Request.Body)
		if err != nil {
			fmt.Println(err)
			//l.tplName="error.html"
			l.Ctx.WriteString("抱歉，用户信息解析失败")
			return
		}
		//fmt.Println("12121212", bodyData)
		//
		var user models.User
		json.Unmarshal(bodyData, &user)
		//err := l.ParseForm(&user)
		//fmt.Println("解析后的user：", user)
		if err != nil {

	}
	//用户登录
	//var user models.User
	//err := l.ParseForm(&user)
	//if err != nil {
	//	//l.tplName="error.html"
	//	l.Ctx.WriteString("抱歉，用户信息解析失败")
	//	return
	//}
	//fmt.Println("电话：",user.UserName,user.Password)

	//	查询数据库的用户信息
	_, err= user.QueryUser()
	if err != nil {
		fmt.Println(err.Error())
		l.Ctx.WriteString("抱歉，用户登录失败，请重试！")
		return
	}
	//1、BlockChainInfo定义
	GetBlockChainInfo,_:=json.Marshal(util.GetBlockChainInfo())
	l.Data["BlockChainInfo"] =string(GetBlockChainInfo)
	//2、BlockCount使用
	BlockCount,_:=json.Marshal(util.GetBlockCount())
	l.Data["BlockCount"]=string(BlockCount)

	//3、BlockHash
	BlockHahsh,_:=json.Marshal(util.GetBlockHash(0))
	l.Data["BlockHash"]=string(BlockHahsh)
	//4、Balances
	Balances,_:=json.Marshal(util.GetBalances())
	l.Data["Balances"]=string(Balances)
	//5、BestblockHash
	BestBlockHash,_:=json.Marshal(util.GetBestBlockHash())
	l.Data["BestBlockHash"]=string(BestBlockHash)
	//6、Difficulty
	GetDifficulty,_:=json.Marshal(util.GetDifficulty())
	l.Data["GetDifficulty"]=string(GetDifficulty)
	//7、Mininginfo 挖矿信息
	GetMininginfo,_:=json.Marshal(util.GetMininginfo())
	l.Data["GetMininginfo"]=string(GetMininginfo)
	//8、Walletinfo 钱包信息
	GetWalletInfo,_:=json.Marshal(util.GetWallenInfo())
	l.Data["GetWalletInfo"]=string(GetWalletInfo)


	//9、VerifyChain:校验位
	VerifyChain,_:=json.Marshal(util.VerifyChain())
	l.Data["VerifyChain"]=string(VerifyChain)
	//10、pruneblockchain：修建区块高度
	Pruneblockchain,_:=json.Marshal(util.Pruneblockchain(0))
	l.Data["Pruneblockchain"]=string(Pruneblockchain)
	//11、GetBlockHeader:获取区块头信息
	GetBlockHeader,_:=json.Marshal(util.GetBlockHeader(util.Hash0))
	l.Data["GetBlockHeader"]=string(GetBlockHeader)
	//12、Getchaintxstats
	Getchaintxstats,_:=json.Marshal(util.GetChaintxstats())
	l.Data["Getchaintxstats"]=string(Getchaintxstats)
	l.TplName = "response.html"
}
//func (l *LoginController)Post()  {
//	l.TplName="test.html"
//}
//func (l *LoginController)Post()  {
//	l.TplName="ajax.html"
//}


