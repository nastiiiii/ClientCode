package Update

import (
	"awesomeProject/Enteties"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func UpdateAccount(account Enteties.Account, accountID int) {
	url := fmt.Sprintf("http://localhost:8080/Bank/api/accounts/", accountID)

	requestBody, err := json.Marshal(account)
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
