package entity

type SoftForks struct {
	Bip34  Bip `json:"bip_34"`
	Bip66  Bip `json:"bip_66"`
	Bip65  Bip `json:"bip_65"`
	Csv    Bip `json:"csv"`
	Segwit Bip `json:"segwit"`
}
