package mails

import (
	"fmt"
	"net/smtp"
)

func EmailNotifier() {
	from := "from@gmail.com"
	password := "<Email Password>"

	to := []string{
		"sender@example.com",
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
