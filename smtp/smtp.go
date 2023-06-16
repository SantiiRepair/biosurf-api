package smtp

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
	"os"

	"github.com/gofiber/fiber/v2"
)

func Mailer(c *fiber.Ctx) error{
	from := mail.Address{Name: "", Address: os.Getenv("MAIL_ADDRESS")}
	password := os.Getenv("MAIL_PASSWORD")
	to := mail.Address{Name: "", Address: c.FormValue("email")}

	subj := "This is the email subject"
	body := "This is an example body.\n With two lines."

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	hostname := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	server := fmt.Sprintf("%s:%s", hostname, port)

	host, _, _ := net.SplitHostPort(server)

	auth := smtp.PlainAuth("", from.Address, password, host)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", server, tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	s, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}

	if err = s.Auth(auth); err != nil {
		log.Panic(err)
	}

	if err = s.Mail(from.Address); err != nil {
		log.Panic(err)
	}

	if err = s.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}

	w, err := s.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	s.Quit()

	return err
}
