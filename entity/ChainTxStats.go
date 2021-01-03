package entity

type ChainTxStats struct {
	Time int `json:"time"`
	Txcount int `json:"txcount"`
	Window_final_block_hash string `json:"window_final_block_hash"`
	Window_final_block_height int `json:"window_final_block_height"`
	Window_block_count int `json:"window_block_count"`
	Window_tx_count int `json:"window_tx_count"`
	Window_interval int `json:"window_interval"`
	Txrate float64 `json:"txrate"`
}
