package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"awesomeProject/Enteties"
	"awesomeProject/Requests/AccountBased"
	"awesomeProject/Requests/Delete"
	"awesomeProject/Requests/Get"
	"awesomeProject/Requests/Post"
	"awesomeProject/Requests/Update"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println()
		fmt.Println("=== Command Line Menu ===")
		fmt.Println("1) Create a Student")
		fmt.Println("2) Create an Account for a Student")
		fmt.Println("3) Get a Student by ID")
		fmt.Println("4) Get an Account by ID")
		fmt.Println("5) Get Accounts by Student ID")
		fmt.Println("6) Update a Student")
		fmt.Println("7) Update an Account")
		fmt.Println("8) Delete a Student")
		fmt.Println("9) Delete an Account")
		fmt.Println("10) Process a Transaction (Deposit/Withdrawal)")
		fmt.Println("11) Process a Transfer between two Accounts")
		fmt.Println("12) Show All Students")
		fmt.Println("13) Show All Accounts")
		fmt.Println("0) Exit")

		fmt.Print("\nEnter your choice: ")
		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)

		switch choiceStr {
		case "1":
			//Tested
			createStudentCLI(reader)
		case "2":
			//Tested
			createAccountForStudentCLI(reader)
		case "3":
			//Tested
			getStudentByIDCLI(reader)
		case "4":
			//Tested
			getAccountByIDCLI(reader)
		case "5":
			//Tested
			getAccountsByStudentIDCLI(reader)
		case "6":
			//Tested
			updateStudentCLI(reader)
		case "7":
			//Tested
			updateAccountCLI(reader)
		case "8":
			//Tested
			deleteStudentCLI(reader)
		case "9":
			//Tested
			deleteAccountCLI(reader)
		case "10":
			processTransactionCLI(reader)
		case "11":
			//TODO
			processTransferCLI(reader)
		case "12":
			//Tested
			showAllStudentsCLI()
		case "13":
			//Tested
			showAllAccountsCLI()
		case "0":
			fmt.Println("Exiting application.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func createStudentCLI(reader *bufio.Reader) {
	fmt.Println("=== Create a Student ===")
	fmt.Print("Enter student name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter student address: ")
	address, _ := reader.ReadString('\n')
	address = strings.TrimSpace(address)

	fmt.Print("Enter student email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Enter student phone: ")
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)

	newStudent := Enteties.Students{
		StudentName:    name,
		StudentAddress: address,
		StudentEmail:   email,
		StudentPhone:   phone,
	}

	// Make the POST request to create a student
	err := Post.CreateStudent(newStudent)
	if err != nil {
		log.Println("Error creating student:", err)
		return
	}

	fmt.Println("Student created successfully!")
}

func createAccountForStudentCLI(reader *bufio.Reader) {
	fmt.Println("=== Create an Account for a Student ===")
	fmt.Print("Enter student ID: ")
	studentIDStr, _ := reader.ReadString('\n')
	studentIDStr = strings.TrimSpace(studentIDStr)
	studentID, err := strconv.Atoi(studentIDStr)
	if err != nil {
		log.Println("Invalid student ID:", err)
		return
	}

	fmt.Print("Enter account alias: ")
	alias, _ := reader.ReadString('\n')
	alias = strings.TrimSpace(alias)

	fmt.Print("Enter initial account balance: ")
	balanceStr, _ := reader.ReadString('\n')
	balanceStr = strings.TrimSpace(balanceStr)
	balance, err := strconv.ParseFloat(balanceStr, 64)
	if err != nil {
		log.Println("Invalid balance:", err)
		return
	}

	newAccount := Enteties.AddAccount{
		AccountAlias:   alias,
		AccountBalance: balance,
	}

	err = Post.CreateAccountForStudent(newAccount, studentID)
	if err != nil {
		log.Println("Error creating account:", err)
		return
	}

	fmt.Println("Account created successfully for student ID:", studentID)
}

func getStudentByIDCLI(reader *bufio.Reader) {
	fmt.Println("=== Get Student by ID ===")
	fmt.Print("Enter student ID: ")
	studIDStr, _ := reader.ReadString('\n')
	studIDStr = strings.TrimSpace(studIDStr)
	studID, err := strconv.Atoi(studIDStr)
	if err != nil {
		log.Println("Invalid student ID:", err)
		return
	}

	student, err := Get.GetByStudentID(studID)
	if err != nil {
		log.Println("Error getting student by ID:", err)
		return
	}

	studentJSON, _ := json.MarshalIndent(student, "", "  ")
	fmt.Println("Student:", string(studentJSON))
}

func getAccountByIDCLI(reader *bufio.Reader) {
	fmt.Println("=== Get Account by ID ===")
	fmt.Print("Enter account ID: ")
	accountIDStr, _ := reader.ReadString('\n')
	accountIDStr = strings.TrimSpace(accountIDStr)
	accountID, err := strconv.Atoi(accountIDStr)
	if err != nil {
		log.Println("Invalid account ID:", err)
		return
	}

	account := Get.GetByAccountID(accountID)

	accJSON, _ := json.MarshalIndent(account, "", "  ")
	fmt.Println("Account:", string(accJSON))
}

func getAccountsByStudentIDCLI(reader *bufio.Reader) {
	fmt.Println("=== Get Accounts by Student ID ===")
	fmt.Print("Enter student ID: ")
	studIDStr, _ := reader.ReadString('\n')
	studIDStr = strings.TrimSpace(studIDStr)
	studID, err := strconv.Atoi(studIDStr)
	if err != nil {
		log.Println("Invalid student ID:", err)
		return
	}

	accounts := Get.GetAccountsByStudentID(studID)

	accJSON, _ := json.MarshalIndent(accounts, "", "  ")
	fmt.Println("Accounts:", string(accJSON))
}

func updateStudentCLI(reader *bufio.Reader) {
	fmt.Println("=== Update Student ===")
	fmt.Print("Enter student ID to update: ")
	studIDStr, _ := reader.ReadString('\n')
	studIDStr = strings.TrimSpace(studIDStr)
	studID, err := strconv.Atoi(studIDStr)
	if err != nil {
		log.Println("Invalid student ID:", err)
		return
	}

	fmt.Print("Enter new student name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter new student address: ")
	address, _ := reader.ReadString('\n')
	address = strings.TrimSpace(address)

	fmt.Print("Enter new student email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Enter new student phone: ")
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)

	updatedStudent := Enteties.Students{
		StudentID:      studID,
		StudentName:    name,
		StudentAddress: address,
		StudentEmail:   email,
		StudentPhone:   phone,
	}

	err = Update.UpdateStudent(updatedStudent)
	if err != nil {
		log.Println("Error updating student:", err)
		return
	}

	fmt.Println("Student updated successfully!")
}

func updateAccountCLI(reader *bufio.Reader) {
	fmt.Println("=== Update Account ===")
	fmt.Print("Enter account ID to update: ")
	accIDStr, _ := reader.ReadString('\n')
	accIDStr = strings.TrimSpace(accIDStr)
	accID, err := strconv.Atoi(accIDStr)
	if err != nil {
		log.Println("Invalid account ID:", err)
		return
	}

	fmt.Print("Enter new student ID (owner of the account): ")
	studIDStr, _ := reader.ReadString('\n')
	studIDStr = strings.TrimSpace(studIDStr)
	studID, err := strconv.Atoi(studIDStr)
	if err != nil {
		log.Println("Invalid student ID:", err)
		return
	}

	fmt.Print("Enter new account alias: ")
	alias, _ := reader.ReadString('\n')
	alias = strings.TrimSpace(alias)

	fmt.Print("Enter new account balance: ")
	balanceStr, _ := reader.ReadString('\n')
	balanceStr = strings.TrimSpace(balanceStr)
	balance, err := strconv.ParseFloat(balanceStr, 64)
	if err != nil {
		log.Println("Invalid balance:", err)
		return
	}

	updatedAccount := Enteties.Account{
		AccountID:      accID,
		StudentID:      studID,
		AccountAlias:   alias,
		AccountBalance: balance,
	}

	err = Update.UpdateAccount(updatedAccount)
	if err != nil {
		log.Println("Error updating account:", err)
		return
	}

	fmt.Println("Account updated successfully!")
}

func deleteStudentCLI(reader *bufio.Reader) {
	fmt.Println("=== Delete Student ===")
	fmt.Print("Enter student ID to delete: ")
	studIDStr, _ := reader.ReadString('\n')
	studIDStr = strings.TrimSpace(studIDStr)
	studID, err := strconv.Atoi(studIDStr)
	if err != nil {
		log.Println("Invalid student ID:", err)
		return
	}

	err = Delete.DeleteStudent(studID)
	if err != nil {
		log.Println("Error deleting student:", err)
		return
	}

	fmt.Println("Student deleted successfully!")
}

func deleteAccountCLI(reader *bufio.Reader) {
	fmt.Println("=== Delete Account ===")
	fmt.Print("Enter account ID to delete: ")
	accIDStr, _ := reader.ReadString('\n')
	accIDStr = strings.TrimSpace(accIDStr)
	accID, err := strconv.Atoi(accIDStr)
	if err != nil {
		log.Println("Invalid account ID:", err)
		return
	}

	err = Delete.DeleteAccount(accID)
	if err != nil {
		log.Println("Error deleting account:", err)
		return
	}

	fmt.Println("Account deleted successfully!")
}

func processTransactionCLI(reader *bufio.Reader) {
	fmt.Println("=== Process a Transaction (Deposit/Withdrawal) ===")

	fmt.Print("Enter account ID: ")
	accIDStr, _ := reader.ReadString('\n')
	accIDStr = strings.TrimSpace(accIDStr)
	accID, err := strconv.Atoi(accIDStr)
	if err != nil {
		log.Println("Invalid account ID:", err)
		return
	}

	fmt.Print("Enter operation type (deposit/withdraw): ")
	transactionType, _ := reader.ReadString('\n')
	transactionType = strings.TrimSpace(transactionType)

	fmt.Print("Enter amount: ")
	amountStr, _ := reader.ReadString('\n')
	amountStr = strings.TrimSpace(amountStr)
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		log.Println("Invalid amount:", err)
		return
	}

	trans := Enteties.Transaction{
		AccountID: accID,
		Operation: transactionType,
		Amount:    amount,
	}

	fmt.Println("\nChoose Transaction Processing Method:")
	fmt.Println("1) ProcessTransaction")
	fmt.Println("2) ProcessTransactionFeature")
	choiceStr, _ := reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr)

	switch choiceStr {
	case "1":
		AccountBased.ProcessTransaction(trans, accID)
		fmt.Println("Transaction processed successfully (legacy method)!")
	case "2":
		err := AccountBased.ProcessTransactionFeature(trans, accID)
		if err != nil {
			log.Println("ProcessTransactionFeature went wrong:", err)
			return
		}
		fmt.Println("Transaction processed successfully (feature method)!")
	default:
		fmt.Println("Invalid choice. Transaction was not processed.")
	}
}

func processTransferCLI(reader *bufio.Reader) {
	fmt.Println("=== Process Transfer between Two Accounts ===")

	fmt.Print("Enter FROM account ID: ")
	fromAccIDStr, _ := reader.ReadString('\n')
	fromAccIDStr = strings.TrimSpace(fromAccIDStr)
	fromAccID, err := strconv.Atoi(fromAccIDStr)
	if err != nil {
		log.Println("Invalid FROM account ID:", err)
		return
	}

	fmt.Print("Enter TO account ID: ")
	toAccIDStr, _ := reader.ReadString('\n')
	toAccIDStr = strings.TrimSpace(toAccIDStr)
	toAccID, err := strconv.Atoi(toAccIDStr)
	if err != nil {
		log.Println("Invalid TO account ID:", err)
		return
	}

	fmt.Print("Enter FROM student ID: ")
	fromStudIDStr, _ := reader.ReadString('\n')
	fromStudIDStr = strings.TrimSpace(fromStudIDStr)
	fromStudID, err := strconv.Atoi(fromStudIDStr)
	if err != nil {
		log.Println("Invalid FROM student ID:", err)
		return
	}

	fmt.Print("Enter TO student ID: ")
	toStudIDStr, _ := reader.ReadString('\n')
	toStudIDStr = strings.TrimSpace(toStudIDStr)
	toStudID, err := strconv.Atoi(toStudIDStr)
	if err != nil {
		log.Println("Invalid TO student ID:", err)
		return
	}

	fmt.Print("Enter amount to transfer: ")
	amountStr, _ := reader.ReadString('\n')
	amountStr = strings.TrimSpace(amountStr)
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		log.Println("Invalid transfer amount:", err)
		return
	}

	transfer := Enteties.Transfer{
		FromAccountID: fromAccID,
		ToAccountID:   toAccID,
		FromStudentID: fromStudID,
		ToStudentID:   toStudID,
		Amount:        amount,
	}
	fmt.Println("\nChoose Transfer Processing Method:")
	fmt.Println("1) ProcessTransfer")
	fmt.Println("2) ProcessTransferFeature")
	choiceStr, _ := reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr)

	switch choiceStr {
	case "1":
		AccountBased.ProcessTransfer(transfer)
		fmt.Println("Transfer processed successfully (legacy method)!")
	case "2":
		err := AccountBased.ProcessTransferFeature(transfer)
		if err != nil {
			log.Fatal("ProcessTransferFeature went wrong:", err)
		}
		fmt.Println("Transfer processed successfully (feature method)!")
	default:
		fmt.Println("Invalid choice. Transfer was not processed.")
	}
}

func showAllStudentsCLI() {
	fmt.Println("=== Show All Students ===")
	students, err := Get.GetAllStudents()
	if err != nil {
		log.Println("Error getting all students:", err)
		return
	}
	studJSON, _ := json.MarshalIndent(students, "", "  ")
	fmt.Println("All Students:", string(studJSON))
}

func showAllAccountsCLI() {
	fmt.Println("=== Show All Accounts ===")
	accounts, err := Get.GetAllAccounts()
	if err != nil {
		log.Println("Error getting all accounts:", err)
		return
	}
	accJSON, _ := json.MarshalIndent(accounts, "", "  ")
	fmt.Println("All Accounts:", string(accJSON))
}
