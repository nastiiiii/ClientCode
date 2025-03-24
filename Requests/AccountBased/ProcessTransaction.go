package AccountBased

import (
	"awesomeProject/Enteties"
	"awesomeProject/Requests/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// modifyAccountBalance (aka withdraw and deposit): PUT /api/accounts/{account_id}
func ProcessTransaction(transaction Enteties.Transaction, accountID int) {
	url := fmt.Sprintf("%vapi/accounts/%d", config.Domain, accountID)

	requestBody, err := json.Marshal(transaction)
	if err != nil {
		fmt.Errorf("error encoding JSON: %w", err)
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Errorf("error creating PUT request %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("error sendiong PUT request %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("error reafing response body %w", err)
	}

	log.Printf("Server response status: %s", resp.Status)
	log.Printf("Server response body: %s", string(responseBody))

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fmt.Errorf("server responded with non-2xx status code: %d", resp.StatusCode)
	}
}
