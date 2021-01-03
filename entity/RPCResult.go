package entity

type RPCResult struct {
	Result interface{} `json:"result"`
	Error  string      `json:"error"`
	Id     int64       `json:"id"`
}
