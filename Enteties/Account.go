package Enteties

type Account struct {
	AccountID      int     `json:"accountID"`
	StudentID      int     `json:"studentID"`
	AccountAlias   string  `json:"accountAlias"`
	AccountBalance float64 `json:"accountBalance"`
}
