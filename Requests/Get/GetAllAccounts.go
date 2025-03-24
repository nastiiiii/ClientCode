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

func GetAllAccounts() {
	url := fmt.Sprintf("%vapi/accounts", config.Domain)

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

	var accounts []Enteties.Account

	err = json.Unmarshal(body, &accounts)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	for _, account := range accounts {
		fmt.Printf("Account: %+v\n", account)
	}
}
