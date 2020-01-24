package main

import (
	"fmt"
	"os"

	"github.com/sfreiberg/gotwilio"
	"github.com/subosito/gotenv"
)

func init() {
	err := gotenv.Load(".env")
	if err != nil {
		fmt.Println("error loading environment", err)
	}
}

func main() {
	accountSid := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	from := os.Getenv("FROM_NUMBER")
	to := os.Getenv("TO_NUMBER")
	message := "Testing Twilio Microservice"
	twilio.SendSMS(from, to, message, "", "")
}
