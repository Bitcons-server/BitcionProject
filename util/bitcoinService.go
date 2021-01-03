package util

import (
	"BitcoinProject/entity"
	"encoding/json"
	"fmt"
	"github.com/mapstructure"
)

/*
*  获取指定高度区块的哈希
 */
func GetBlockHash(height int) interface{} {
	reqBytes := PrepareJson(GETBLOCKHASH, []interface{}{height})
	body := Dopost(RPCURL, reqBytes)
	RPCResult := GetResult(body)
	return RPCResult.Result
}
/*
*  被修剪的区块的高度
*/
func Pruneblockchain(height int) interface{} {
	reqBytes:=PrepareJson(PRUNEBLOCKCHAIN,[]interface{}{height})
	body:=Dopost(RPCURL,reqBytes)
	RPCResult:=GetResult(body)
	return  RPCResult.Result
}
/*
* 获取区块头的信息
*/
func GetBlockHeader(hash string)*entity.BlockHeader  {
	reqBytes:=PrepareJson(GETBLOCKHEADER,[]interface{}{hash})
	body:=Dopost(RPCURL,reqBytes)
	RPCResult:=GetResult(body)
	jsonstr,err:=json.Marshal(RPCResult.Result)
	if err!=nil{
		return nil
	}
	BlockHeader:=new(entity.BlockHeader)
	err=json.Unmarshal(jsonstr,BlockHeader)
	if err!=nil{
		fmt.Println(err.Error())
		return nil
	}
	return BlockHeader
}

/*
* 获取区块链当前状态
 */
func GetBlockChainInfo() *entity.BlockChainInfo {
	reqBytes := PrepareJson(GETBLOCKCHAININFO, nil)
	body := Dopost(RPCURL, reqBytes)
	RPCresult := GetResult(body)
	jsonstr, err := json.Marshal(RPCresult.Result)
	if err != nil {
		return nil
	}
	BlochChain := new(entity.BlockChainInfo)
	err = json.Unmarshal(jsonstr, BlochChain)
	//fmt.Println(jsonstr)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return BlochChain
}

/*
*  查询钱包余额
 */
func GetBalances() *entity.Balances {
	reqBytes := PrepareJson(GETBALANCES, nil)
	body := Dopost(RPCURL, reqBytes)
	RPCresult := GetResult(body)
	jsonstr, err := json.Marshal(RPCresult.Result)
	if err != nil {
		return nil
	}
	Balances := new(entity.Balances)
	err = json.Unmarshal(jsonstr, Balances)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return Balances
}

/*
* 获取挖矿信息
 */
func GetMininginfo() *entity.Mininginfo {
	reqBytes := PrepareJson(GETMININGINFO, nil)
	body := Dopost(RPCURL, reqBytes)
	RPCresult := GetResult(body)
	jsonstr, err := json.Marshal(RPCresult.Result)
	if err != nil {
		return nil
	}
	Mininginfo := new(entity.Mininginfo)
	err = json.Unmarshal(jsonstr, Mininginfo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return Mininginfo
}
/*
* 获取钱包信息
*/
func GetWallenInfo()*entity.WalletInfo  {
	reqBytes:=PrepareJson(GETWALLETINFO,nil)
	body:=Dopost(RPCURL,reqBytes)
	RPCresult:=GetResult(body)
	jsonstr,err:=json.Marshal(RPCresult.Result)
	if err!=nil{
		return nil
	}
	WalletInfo:=new(entity.WalletInfo)
	err =json.Unmarshal(jsonstr,WalletInfo)
	if err!=nil{
		fmt.Println(err.Error())
		return nil
	}
	return WalletInfo
}

/*
* 获取区块数量
 */
func GetBlockCount() interface{} {
	reqBytes := PrepareJson(GETBLOCKCOUNT, nil)
	body := Dopost(RPCURL, reqBytes)
	RPCResult := GetResult(body)
	return RPCResult.Result
}

/*
* 获取链头区块哈希
 */
func GetBestBlockHash() interface{} {
	reqBytes := PrepareJson(GETBESTBLOCKHASH, nil)
	body := Dopost(RPCURL, reqBytes)
	RPCResult := GetResult(body)
	return RPCResult.Result
}
/*
* 校验本地区块
*/
func VerifyChain()interface{}  {
	reqBytes:=PrepareJson(VERIFYCHAIN,nil)
	body:=Dopost(RPCURL,reqBytes)
	RPCResult:=GetResult(body)
	return RPCResult.Result
}

/*
* 获取出块难度
 */
func GetDifficulty() float64{
	reqBytes:=PrepareJson(GETDIFFICULTY,nil)
	body:=Dopost(RPCURL,reqBytes)
	RPCResult:=GetResult(body)
	return RPCResult.Result.(float64)
}
//获取区块交易状态
func GetChaintxstats() *entity.ChainTxStats {
	reqBytes:=PrepareJson(GETCHAINTXSTATS,nil)
	body:=Dopost(RPCURL,reqBytes)
	RPCResult:=GetResult(body)
	jsonstr,err:=json.Marshal(RPCResult.Result)
	if err!=nil{
		return nil
	}
	ChainTxStats:=new(entity.ChainTxStats)
	err =json.Unmarshal(jsonstr,ChainTxStats)
	if err!=nil{
		fmt.Println(err.Error())
		return nil
	}
	return ChainTxStats
}
//
func GetChainTips() []*entity.Chaintips  {
	reqBytes:=PrepareJson(GETCHAINTIPS,nil)
	body:=Dopost(RPCURL,reqBytes)
	RPCresult:=GetResult(body)
	//jsonstr,err:=json.Marshal(RPCresult.Result)
	//if err!=nil{
	//	return nil
	//}
	fmt.Println(RPCresult)
	var chaintips []*entity.Chaintips
	err:=mapstructure.WeakDecode(RPCresult,&chaintips)
	if err != nil {
		fmt.Println(err.Error())
	}
	//err =json.Unmarshal(jsonstr,Chaintips)
	//if err!=nil{
	//	fmt.Println(err.Error())
	//	return nil
	//}
	return chaintips
}

//func GetBlockHash(height int) interface{} {
//	reqBytes:=PrepareJson(GETBLOCKHASH,[]interface{}{height})
//	body:=Dopost(RPCURL,reqBytes)
//	RPCResult:=GetResult(body)
//	return RPCResult.Result
//
//}
