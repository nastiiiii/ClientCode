package main

import (
	"awesomeProject/Enteties"
	"awesomeProject/Requests/Delete"
	"awesomeProject/Requests/Get"
	"awesomeProject/Requests/Post"
	"awesomeProject/Requests/Update"
)

// "http://loaclhost/Bank/api/students"
func main() {
	Get.GetByAccountID(1)
	Get.GetByStudentID(1)
	Get.GetAccountsByStudentID(1)
	Get.GetAllAccounts()
	Get.GetAllStudents()

	acc := Enteties.Account{
		StudentID:      1,
		AccountAlias:   "TestAccount",
		AccountBalance: 170,
	}
	stud := Enteties.Students{
		StudentName:    "Name",
		StudentAddress: "TestAddress",
		StudentEmail:   "nas.nas@uwu.com",
		StudentPhone:   "+5554623982",
	}
	Post.CreateAccountForStudent(acc, 1)
	Post.CreateStudent(stud)

	stud.StudentName = "NewName"
	stud.StudentEmail = "NewEmail"
	stud.StudentPhone = "232323"

	acc.AccountBalance = 23

	Update.UpdateStudent(stud, 1)
	Update.UpdateAccount(acc, 1)

	Delete.DeleteAccount(3)
	Delete.DeleteStudent(3)
}
