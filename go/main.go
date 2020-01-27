package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/sfreiberg/gotwilio"
	"github.com/subosito/gotenv"
)

type textRequestStruct struct {
	From string
	To   string
	Body string
}

func init() {
	err := gotenv.Load(".env")
	if err != nil {
		fmt.Println("error loading environment", err)
	}
}

func main() {
	log.Println("Twilio message sender up and listening on port 8087")
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/text", textHandler)
	log.Fatal(http.ListenAndServe(":8087", nil))
}

func healthz(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "Hello, %s!", req.URL.Path[1:])
}

func textHandler(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var textRequest textRequestStruct
	err = json.Unmarshal(body, &textRequest)
	if err != nil {
		panic(err)
	}
	log.Println(textRequest.From, textRequest.To, textRequest.Body)
	//TODO add better error handling
	sendText(textRequest.From, textRequest.To, textRequest.Body)
}

func sendText(from string, to string, body string) {
	accountSid := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)
	//TODO Sanitize from, to, and body
	twilio.SendSMS(from, to, body, "", "")
}
