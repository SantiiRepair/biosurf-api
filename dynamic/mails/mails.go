package mails

import (
	"fmt"
	"net/smtp"
	"os"
)

func HandleNotifier(HandleNotifierring) {
	from := os.Getenv("MAIL_ADDRESS")
	password := os.Getenv("MAIL_PASSWORD")

	to := []string{
		recipient,
	}

	host := "smtp.gmail.com"
	port := "465"

	message := []byte("This is a test email message.")

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(host+":"+port, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
}package mails

import (
	fiber "github.com/gofiber/fiber/v2"
)

func EmailNotifier(r *fiber.App) {
	r.Post("/notifier", HandleReport)
}
