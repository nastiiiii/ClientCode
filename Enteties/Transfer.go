package Enteties

type Transfer struct {
	FromAccountID int     `json:"fromAccountID"`
	ToAccountID   int     `json:"toAccountID"`
	FromStudentID int     `json:"fromStudentID"`
	ToStudentID   int     `json:"toStudentID"`
	Amount        float64 `json:"amount"`
}
