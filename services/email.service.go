package service

import (
	form "employee/app/interface"
	"net/smtp"
	"os"

	"github.com/sirupsen/logrus"
)

func SendEmail(user *form.User) {
	from := os.Getenv("EMAIL_DEFAULT")
	password := os.Getenv("EMAIL_PASSWORD")
	to := []string{user.Email}
	smtpHost := os.Getenv("EMAIL_HOST")
	smtpPort := os.Getenv("EMAIL_PORT")

	// message := []byte("Hello <b>" + user.Email + "</b><br><p> Your password is: <b>" + user.Password + "</b></p>")
	message := []byte("To: " + user.Email + "\r\n" +
		"Subject: Create new account\r\n" +
		"\r\n" +
		"Hello " + user.Email + "\r\nYour password is: " + user.Password + "\r\n")

	// Create authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send actual message
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Print("Email sent successfully")
}
