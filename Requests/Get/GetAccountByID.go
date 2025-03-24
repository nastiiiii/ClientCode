package Get

import (
	"awesomeProject/Enteties"
	"awesomeProject/Requests/config"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// GetByAccountID retrieves an account by its ID and parses the response JSON into an Account struct.
func GetByAccountID(id int) Enteties.Account {
	url := fmt.Sprintf("%vapi/accounts/%d", config.Domain, id)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error performing GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Received non-OK response: %d %s", resp.StatusCode, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var account Enteties.Account

	err = json.Unmarshal(body, &account)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	log.Printf("Account Parsed: %+v", account)
	return account
}
