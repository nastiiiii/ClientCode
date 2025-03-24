package Enteties

type Transaction struct {
	AccountID   int `json:"accountID"`
	Transaction int
	Amount      int `json:"amount"`
}
