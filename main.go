package main

import (
	"BitcoinProject/db_mysql"
	_ "BitcoinProject/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db_mysql.ConnectDB()
	//fmt.Println("第0个区块的hash为：", util.GetBlockHash(0))
	//fmt.Println("",util.GetBlockChainInfo().Verificationprogress)
	//fmt.Println("当前的钱包余额是：",util.GetBalances().Mine.Trusted)
	//fmt.Println("当前的区块数量为：",util.GetBlockCount())
	//fmt.Println("",util.GetBestBlockHash())

	//fmt.Println("",util.GetDifficulty())
	//fmt.Println("",util.GetMininginfo().Blocks)
	//fmt.Println("",util.GetWallenInfo())
	//fmt.Println("",util.VerifyChain())
	//fmt.Println("",util.Pruneblockchain(0))
	//fmt.Println("",util.GetBlockHeader(util.Hash0).Version)
	//fmt.Println("",util.GetChainTips()[0])
	//fmt.Println("",util.GetChaintxstats().Time)
	//静态资源路径设置
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/image", "./static/image")

	beego.Run(":8090")

}

