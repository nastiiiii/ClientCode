package Delete

import (
	"awesomeProject/Requests/config"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func DeleteStudent(studentID int) {
	url := fmt.Sprintf("%vapi/students/delete/%d", config.Domain, studentID)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		fmt.Errorf("error creating DELETE request: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("error sending DELETE request: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("error reading response body: %w", err)
	}

	log.Printf("Server response status: %s", resp.Status)
	log.Printf("Server response body: %s", string(responseBody))

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fmt.Errorf("server responded with non-2xx status: %d", resp.StatusCode)
	}
}
