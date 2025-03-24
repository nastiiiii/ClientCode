package Get

import (
	"awesomeProject/Requests/config"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetAccountsByStudentID(id int) {
	url := fmt.Sprintf("%vapi/accounts/studentID/%d", config.Domain, id)
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
