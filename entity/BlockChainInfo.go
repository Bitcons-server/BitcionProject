package entity

type BlockChainInfo struct {
	Chain                string    `json:"chain"`
	Blocks               int64     `json:"blocks"`
	Headers              int64     `json:"headers"`
	Bestblockhash        string    `json:"bestblockhash"`
	Difficulty           float64   `json:"difficulty"`
	Verificationprogress float64   `json:"verificationprogress"`
	Initialblockdownload bool      `json:"initialblockdownload"`
	Chainwork            string    `json:"chainwork"`
	Size_on_disk         int64     `json:"size_on_disk"`
	Pruned               bool      `json:"pruned"`
	Pruneheight          int       `json:"pruneheight"`
	Automatic_pruning    bool      `json:"automatic_pruning"`
	Prune_target_size    int64     `json:"prune_target_size"`
	Softforks            SoftForks `json:"softforks"`
	Warnings             string    `json:"warnings"`
}
