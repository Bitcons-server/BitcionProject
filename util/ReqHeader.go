package util

/*
* 区块链的头部信息
*/
func ReqHeader() map[string]string {
	header := make(map[string]string)
	header["Encoding"] = "UTF-8"
	header["Content-Type"] = "application/json"
	header["Authorization"] = "Basic " + Base64EncodeString(RPCUSER+":"+RPCPASSWORD)
	return header
}
