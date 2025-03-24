package Delete

import (
	"awesomeProject/Requests/config"
	_ "awesomeProject/Requests/config"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func DeleteAccount(accountID int) error {
	url := fmt.Sprintf("%vapi/accounts/delete/%d", config.Domain, accountID)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		fmt.Errorf("error creating DELETE request: %w", err)
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("error sending DELETE request: %w", err)
		return err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("error reading response body: %w", err)
		return err
	}

	log.Printf("Server response status: %s", resp.Status)
	log.Printf("Server response body: %s", string(responseBody))

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fmt.Errorf("server responded with non-2xx status: %d", resp.StatusCode)
		return err
	}
	return nil
}
