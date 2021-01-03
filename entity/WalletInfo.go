package entity

type WalletInfo struct {
	Walletname              string  `json:"walletname"`
	Walletversion           int     `json:"walletversion"`
	Balance                 float64 `json:"balance"`
	Unconfirmed_balance     float64 `json:"unconfirmed_balance"`
	Immature_balance        float64 `json:"immature_balance"`
	Txcount                 int     `json:"txcount"`
	Keypoololdest           int     `json:"keypoololdest"`
	Keypoolsize             int     `json:"keypoolsize"`
	Hdseedid                string  `json:"hdseedid"`
	Keypoolsize_hd_internal int     `json:"keypoolsize_hd_internal"`
	Paytxfee                float64 `json:"paytxfee"`
	Private_keys_enabled    bool    `json:"private_keys_enabled"`
	Avoid_reuse             bool    `json:"avoid_reuse"`
	Scanning                bool    `json:"scanning"`
}
