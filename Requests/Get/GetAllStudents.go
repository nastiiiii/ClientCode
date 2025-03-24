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

func GetAllStudents() {
	url := fmt.Sprintf("%vapi/students", config.Domain)

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

	var students []Enteties.Students

	err = json.Unmarshal(body, &students)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	for _, student := range students {
		fmt.Printf("Student: %+v\n", student)
	}
}
