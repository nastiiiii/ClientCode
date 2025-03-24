package Enteties

type Transaction struct {
	AccountID int     `json:"accountID"`
	Operation string  `json:"operation"`
	Amount    float64 `json:"amount"`
}
