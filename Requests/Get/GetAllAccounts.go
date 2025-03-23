package Get

import (
	"io"
	"log"
	"net/http"
)

func GetAllAccounts() {
	resp, err := http.Get("http://loaclhost/Bank/api/accounts")
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
