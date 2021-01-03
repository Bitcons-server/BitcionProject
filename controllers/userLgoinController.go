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
		fmt.Println("12121212", bodyData)
		//
		var user models.User
		json.Unmarshal(bodyData, &user)
		//err := l.ParseForm(&user)
		fmt.Println("解析后的user：", user)
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
	//BlockChainInfo定义
	l.Data["BlockChainInfo"] =util.GetBlockChainInfo().Difficulty
	l.Data["BlockChainInfo1"] =util.GetBlockChainInfo().Chain
	l.Data["BlockChainInfo2"] =util.GetBlockChainInfo().Blocks
	l.Data["BlockChainInfo3"] =util.GetBlockChainInfo().Headers
	l.Data["BlockChainInfo4"] =util.GetBlockChainInfo().Bestblockhash
	l.Data["BlockChainInfo5"] =util.GetBlockChainInfo().Verificationprogress
	l.Data["BlockChainInfo6"] =util.GetBlockChainInfo().Chainwork
	l.Data["BlockChainInfo7"] =util.GetBlockChainInfo().Pruned
	l.Data["BlockChainInfo8"] =util.GetBlockChainInfo().Softforks
	//BlockCount使用
	l.Data["BlockCount"]=util.GetBlockCount()
	//BlockHash
	l.Data["BlockHash"]=util.GetBlockHash(0)
	//Balances
	l.Data["Balances"]=util.GetBalances().Mine.Trusted
	l.Data["Balances1"]=util.GetBalances().Mine.Immature
	l.Data["Balances2"]=util.GetBalances().Mine.Untrusted_pending
	//BestblockHash
	l.Data["BestBlockHash"]=util.GetBestBlockHash()
	//Difficulty
	l.Data["GetDifficulty"]=util.GetDifficulty()
	//Mininginfo 挖矿信息
	l.Data["GetMininginfo"]=util.GetMininginfo().Blocks
	l.Data["GetMininginfo1"]=util.GetMininginfo().Chain
	l.Data["GetMininginfo2"]=util.GetMininginfo().Difficulty
	l.Data["GetMininginfo3"]=util.GetMininginfo().Warnings
	l.Data["GetMininginfo4"]=util.GetMininginfo().Networkhashps
	l.Data["GetMininginfo5"]=util.GetMininginfo().Pooledtx
	//Walletinfo 钱包信息
	l.Data["GetWalletInfo"]=util.GetWallenInfo().Walletversion
	l.Data["GetWalletInfo1"]=util.GetWallenInfo().Balance
	l.Data["GetWalletInfo2"]=util.GetWallenInfo().Txcount
	l.Data["GetWalletInfo3"]=util.GetWallenInfo().Keypoololdest
	l.Data["GetWalletInfo4"]=util.GetWallenInfo().Keypoolsize
	l.Data["GetWalletInfo5"]=util.GetWallenInfo().Unconfirmed_balance
	l.TplName = "response.html"
	//VerifyChain:校验位
	l.Data["VerifyChain"]=util.VerifyChain()
	//pruneblockchain：修建区块高度
	l.Data["Pruneblockchain"]=util.Pruneblockchain(0)
	//GetBlockHeader:获取区块头信息
	l.Data["GetBlockHeader"]=util.GetBlockHeader(util.Hash0).Height
	l.Data["GetBlockHeader1"]=util.GetBlockHeader(util.Hash0).Difficulty
	l.Data["GetBlockHeader2"]=util.GetBlockHeader(util.Hash0).Chainwork
	l.Data["GetBlockHeader3"]=util.GetBlockHeader(util.Hash0).Bits
	l.Data["GetBlockHeader4"]=util.GetBlockHeader(util.Hash0).Confirmations
	l.Data["GetBlockHeader5"]=util.GetBlockHeader(util.Hash0).Nextblockhash
	l.Data["GetBlockHeader6"]=util.GetBlockHeader(util.Hash0).Time
	l.Data["GetBlockHeader7"]=util.GetBlockHeader(util.Hash0).NTx
	l.Data["GetBlockHeader8"]=util.GetBlockHeader(util.Hash0).VersionHex
	l.Data["GetBlockHeader9"]=util.GetBlockHeader(util.Hash0).Version
	//Getchaintxstats
	l.Data["Getchaintxstats"]=util.GetChaintxstats().Time
	l.Data["Getchaintxstats1"]=util.GetChaintxstats().Txcount
	l.Data["Getchaintxstats2"]=util.GetChaintxstats().Txrate
	l.Data["Getchaintxstats3"]=util.GetChaintxstats().Window_block_count
	l.Data["Getchaintxstats4"]=util.GetChaintxstats().Window_final_block_hash
	l.Data["Getchaintxstats5"]=util.GetChaintxstats().Window_final_block_height
	l.Data["Getchaintxstats6"]=util.GetChaintxstats().Window_interval
	l.Data["Getchaintxstats7"]=util.GetChaintxstats().Window_tx_count
}
//func (l *LoginController)Post()  {
//	l.TplName="test.html"
//}


