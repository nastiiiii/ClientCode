package Post

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

func CreateAccountForStudent(account Enteties.AddAccount, studentID int) error {
	postBody, err := json.Marshal(account)
	if err != nil {
		log.Fatal("Error encoding JSON %v", err)
		return err
	}

	requestBody := bytes.NewBuffer(postBody)

	url := fmt.Sprintf("%vapi/accounts/studentID/%d", config.Domain, studentID)

	resp, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		log.Fatal("Error making POST request: %v", err)
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error handling response body %v", err)
		return err
	}

	responseString := string(body)
	fmt.Println("Response Body", responseString)
	return nil
}
