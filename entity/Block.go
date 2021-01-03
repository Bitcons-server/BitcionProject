package entity

type Block struct {
	Hash string `json:"hash"`
	Confirmations int `json:"confirmations"`
	Srippedsize int `json:"strippedsize"`
	Size int `json:"size"`
	Weight int `json:"weight"`
	Version int `json:"version"`
	VersionHex string `json:"version_hex"`
	Merkleroot string `json:"merkleroot"`
	Height int `json:"height"`
	Time int `json:"time"`
	Mediantime int `json:"mediantime"`
	Nonce int `json:"nonce"`
	Bits string `json:"bits"`
	Difficulty int `json:"difficulty"`
	Chainwork string `json:"chainwork"`
	NTx int `json:"nTx"`
}