package entity

type BlockHeader struct {
	Hash          string `json:"hash"`
	Confirmations int    `json:"confirmations"`
	Height        int    `json:"height"`
	Version       int    `json:"version"`
	VersionHex    string `json:"version_hex"`
	Merkleroot    string `json:"merkleroot"`
	Time          int    `json:"time"`
	Mediantime    int    `json:"mediantime"`
	Nonce         int    `json:"nonce"`
	Bits          string `json:"bits"`
	Difficulty    int    `json:"difficulty"`
	Chainwork     string `json:"chainwork"`
	NTx           int    `json:"n_tx"`
	Nextblockhash string `json:"nextblockhash"`
}
