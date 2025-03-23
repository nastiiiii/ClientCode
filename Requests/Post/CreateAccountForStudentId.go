package Post

import (
	"awesomeProject/Enteties"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateAccountForStudent(account Enteties.Account) {
	postBody, err := json.Marshal(account)
	if err != nil {
		log.Fatal("Error encoding JSON %v", err)
	}

	requestBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("http://localhost:8080/Bank/api/students", "application/json", requestBody)
	if err != nil {
		log.Fatal("Error making POST request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error handling response body %v", err)
	}

	responseString := string(body)
	fmt.Println("Response Body", responseString)
}
