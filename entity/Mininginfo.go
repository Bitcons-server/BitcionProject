package entity

type Mininginfo struct {
	Blocks int `json:"blocks"`
	Difficulty float64 `json:"difficulty"`
	Networkhashps float32`json:"networkhashps"`
	Pooledtx int `json:"pooledtx"`
	Chain string `json:"chain"`
	Warnings  string `json:"warnings"`
}
