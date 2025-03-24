package AccountBased

import (
	"awesomeProject/Enteties"
	"awesomeProject/Requests/Get"
	"fmt"
)

func ProcessTransferFeature(transfer Enteties.Transfer) error {
	fromAccounts := Get.GetAccountsByStudentID(transfer.FromStudentID)
	if len(fromAccounts) == 0 {
		return fmt.Errorf("no accounts found for student (fromStudentID=%d)", transfer.FromStudentID)
	}

	startIndex := -1
	for i, acct := range fromAccounts {
		if acct.AccountID == transfer.FromAccountID {
			startIndex = i
			break
		}
	}
	if startIndex == -1 {
		return fmt.Errorf("initial fromAccountID=%d not found among the student's accounts", transfer.FromAccountID)
	}

	amountNeeded := transfer.Amount

	toAccountID := transfer.ToAccountID
	toStudentID := transfer.ToStudentID

	for i := startIndex; i < len(fromAccounts) && amountNeeded > 0; i++ {
		currentAcct := fromAccounts[i]
		if currentAcct.AccountBalance <= 0 {
			continue
		}

		contribution := currentAcct.AccountBalance
		if amountNeeded < contribution {
			contribution = amountNeeded
		}

		partialTransfer := Enteties.Transfer{
			FromAccountID: currentAcct.AccountID,
			ToAccountID:   toAccountID,
			FromStudentID: currentAcct.StudentID,
			ToStudentID:   toStudentID,
			Amount:        contribution,
		}

		ProcessTransfer(partialTransfer)

		amountNeeded -= contribution
	}

	if amountNeeded > 0 {
		return fmt.Errorf("not enough funds to complete transfer of %v from student %d to account %d",
			transfer.Amount, transfer.FromStudentID, transfer.ToAccountID)
	}

	return nil
}
