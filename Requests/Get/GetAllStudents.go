package Get

import (
	"io"
	"log"
	"net/http"
)

func GetAllStudents() {
	resp, err := http.Get("http://loaclhost/Bank/api/students")
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
