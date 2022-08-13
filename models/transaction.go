package models

type Transaction struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Value    float64 `json:"value"`
	Data     string  `json:"data"`
	Type     string  `json:"type"`
	Conta_ID int     `json:"conta_id"`
}
