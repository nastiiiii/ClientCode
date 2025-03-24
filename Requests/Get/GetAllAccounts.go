package Get

import (
	"awesomeProject/Requests/config"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetAllAccounts() {
	url := fmt.Sprintf("%vapi/accounts", config.Domain)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	sb := string(body)
	log.Printf(sb)
}
