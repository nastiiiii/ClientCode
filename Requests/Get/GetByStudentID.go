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

// GetByStudentID retrieves a student by ID and prints their data.
func GetByStudentID(id int) {
	url := fmt.Sprintf("%vapi/students/%d", config.Domain, id)
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

	var student Enteties.Students

	err = json.Unmarshal(body, &student)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	log.Printf("Student: %+v\n", student)
}
