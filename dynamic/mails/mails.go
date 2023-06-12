package mails

import (
	"fmt"
	"net/smtp"
	"os"
)

func EmailNotifier(recipient string) {
	from := os.Getenv("MAIL_ADDRESS")
	password := os.Getenv("MAIL_PASSWORD")

	to := []string{
		recipient,
	}

	host := "smtp.gmail.com"
	port := "587"

	message := []byte("This is a test email message.")

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(host+":"+port, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
}
