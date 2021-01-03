package entity

type Mine struct {
	Trusted  int64 `json:"trusted"`
	Untrusted_pending int64 `json:"untrusted_pending"`
	Immature int64 `json:"immature"`
}
