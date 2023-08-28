package middlewares

import (
	"bytes"
	"encoding/json"
	"net/http"
	"fmt"
	"math"
	"math/rand"
	"os"
)

const apiURL = "https://api.d7networks.com/messages/v1/send"

type Message struct {
	Channel     string   `json:"channel"`
	Recipients  []string `json:"recipients"`
	Content     string   `json:"content"`
	MsgType     string   `json:"msg_type"`
	DataCoding  string   `json:"data_coding"`
}

type MessageGlobals struct {
	Originator string `json:"originator"`
	ReportURL  string `json:"report_url"`
}

type SMSRequest struct {
	Messages       []Message      `json:"messages"`
	MessageGlobals MessageGlobals `json:"message_globals"`
}

func SendSMS(phoneNumber string, otp string) error {
	message := Message{
		Channel:    "sms",
		Recipients: []string{phoneNumber},
		Content:    "Your OTP is: " + otp,
		MsgType:    "text",
		DataCoding: "text",
	}

	messageGlobals := MessageGlobals{
		Originator: "SignOTP",
		ReportURL:  "https://the_url_to_recieve_delivery_report.com",
	}

	smsRequest := SMSRequest{
		Messages:       []Message{message},
		MessageGlobals: messageGlobals,
	}

	data, err := json.Marshal(smsRequest)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+ os.Getenv("D7NETWORKS_BEARER_TOKEN"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func GenerateOTP(num int) string {
	max := int(math.Pow10(num)) - 1
	min := int(math.Pow10(num - 1))
	return fmt.Sprintf("%0*d", num, rand.Intn(max-min+1)+min)
}