package Get

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetByAccountID(id int) {
	url := fmt.Sprintf("http://localhost/Bank/api/accounts/%d", id)
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
