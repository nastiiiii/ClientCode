package Enteties

type Transaction struct {
	AccountID   int `json:"accountID"`
	Transaction string
	Amount      float64 `json:"amount"`
}
