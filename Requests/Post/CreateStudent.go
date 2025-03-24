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

func CreateStudent(student Enteties.Students) {
	postBody, err := json.Marshal(student)
	if err != nil {
		log.Fatal("Error encoding JSON %v", err)
	}

	requestBody := bytes.NewBuffer(postBody)

	url := fmt.Sprintf("%vapi/students", config.Domain)

	resp, err := http.Post(url, "application/json", requestBody)
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
