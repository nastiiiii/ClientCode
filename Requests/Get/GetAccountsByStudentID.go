package Get

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetAccountsByStudentID(id int) {
	url := fmt.Sprintf("http://localhost/Bank/api/accounts/studentID/%d", id)
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
