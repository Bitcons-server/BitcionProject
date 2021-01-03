package entity

type Bip struct {
	Type   string `json:"type"`
	Active bool   `json:"active"`
	Height int    `json:"height"`
}

