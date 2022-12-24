// mail.go contains fucntions for sending of mail

package main

import (
	"log"
	"strconv"

	"gopkg.in/gomail.v2"
)

func sendMail(to string, subject string, body string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", SMTP_ACCOUNT)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	port, err := strconv.ParseInt(SMTP_PORT, 10, 64)
	if err != nil {
		log.Fatal("Cannot parse int from SMTP_PORT")
		return err
	}

	dialer := gomail.NewDialer(SMTP_HOST, int(port), SMTP_ACCOUNT, SMTP_PASSWORD)
	err = dialer.DialAndSend(msg)
	if err != nil {
		log.Fatal("Cannot dial or send message")
		return err
	}
	return nil
}
