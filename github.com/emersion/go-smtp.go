package main

import (
	"log"
	"strings"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
	"github.com/vcgo/test"
)

func main() {
	username := test.Config.Get("email.username").(string)
	password := test.Config.Get("email.password").(string)
	dst0 := test.Config.Get("email.dst0").(string)
	dst1 := test.Config.Get("email.dst1").(string)
	// Set up authentication information.
	auth := sasl.NewPlainClient("", username, password)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{dst0, dst1}
	msg := strings.NewReader("To: " + dst0 + "\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail("smtp.163.com:25", auth, "伟龙的自动提醒", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
