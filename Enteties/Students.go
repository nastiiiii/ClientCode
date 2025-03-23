package Enteties

type Students struct {
	StudentID      int    `json:"studentID"`
	StudentName    string `json:"studentName"`
	StudentAddress string `json:"studentAddress"`
	StudentEmail   string `json:"studentEmail"`
	StudentPhone   string `json:"studentPhone"`
}
