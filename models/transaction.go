package models

type Transaction struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Value    float64 `json:"value"`
	Date     string  `json:"date"`
	Type     string  `json:"type"`
	Conta_ID int     `json:"conta_id"`
}
