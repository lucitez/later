package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

var accountSid = "AC3be74b325d43a3f045a1c8f38d0b0fdf"
var accountToken = "23021642ed3eba51fe74aef72721956b"

func SendSMS(recipient string, content string) error {
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	msgData := url.Values{}

	msgData.Set("To", recipient)
	msgData.Set("From", "12055393449")
	msgData.Set("Body", content)

	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, accountToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make HTTP POST request and return message SID
	resp, _ := client.Do(req)

	var data map[string]interface{}
	decoder := json.NewDecoder(resp.Body)
	err := decoder.Decode(&data)
	fmt.Println(resp.StatusCode)
	fmt.Printf("%v", data)
	return err
}
