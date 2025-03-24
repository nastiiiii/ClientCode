package AccountBased

import (
	"awesomeProject/Enteties"
	"awesomeProject/Requests/Get"
	"fmt"
)

func ProcessTransactionFeature(transaction Enteties.Transaction, accountID int) error {
	currentAccount := Get.GetByAccountID(accountID)

	accounts := Get.GetAccountsByStudentID(currentAccount.StudentID)
	if len(accounts) == 0 {
		return fmt.Errorf("no accounts found for student %d", currentAccount.StudentID)
	}

	// Find index of the current account
	startIndex := -1
	for i, acct := range accounts {
		if acct.AccountID == currentAccount.AccountID {
			startIndex = i
			break
		}
	}
	if startIndex == -1 {
		return fmt.Errorf("current account not found in student's account list")
	}

	amountNeeded := transaction.Amount

	// Iterate from startIndex onward until the amountNeeded is covered or we run out of accounts
	for i := startIndex; i < len(accounts) && amountNeeded > 0; i++ {
		account := accounts[i]

		//how much this account can contribute
		contribution := account.AccountBalance
		if amountNeeded < contribution {
			contribution = amountNeeded
		}

		// Make a transaction for that contribution
		partialTx := Enteties.Transaction{
			AccountID:   account.AccountID,
			Transaction: "deposit",
			Amount:      contribution,
		}
		ProcessTransaction(partialTx, account.AccountID)

		// Deduct from rest
		amountNeeded -= contribution
	}

	if amountNeeded > 0 {
		return fmt.Errorf("insufficient funds across all fallback accounts for total of %v", transaction.Amount)
	}

	return nil
}
