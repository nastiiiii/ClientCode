package Update

import (
	"awesomeProject/Enteties"
	"awesomeProject/Requests/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func UpdateStudent(student Enteties.Students) error {
	url := fmt.Sprintf("%vapi/students", config.Domain)

	requestBody, err := json.Marshal(student)
	if err != nil {
		fmt.Errorf("error encoding JSON: %w", err)
		return err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Errorf("error creating PUT request %w", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("error sendiong PUT request %w", err)
		return err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("error reafing response body %w", err)
		return err
	}

	log.Printf("Server response status: %s", resp.Status)
	log.Printf("Server response body: %s", string(responseBody))

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fmt.Errorf("server responded with non-2xx status code: %d", resp.StatusCode)
		return err
	}
	return nil
}
