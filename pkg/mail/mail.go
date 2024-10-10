package mail

import (
	"crypto/tls"
	"fmt"
	"os"
	"strconv"

	"gopkg.in/mail.v2"
)

func SendEmail(userName, userEmail, userPhone, userMessage string) error {
	senderEmail := os.Getenv("MAIL_USERNAME")
	senderPassword := os.Getenv("MAIL_PASSWORD")
	mailServer := os.Getenv("MAIL_SERVER")
	mailPort := os.Getenv("MAIL_PORT")
	recipientEmail := os.Getenv("NINEPRO_EMAIL")
	useTLS := os.Getenv("MAIL_USE_TLS") == "true" // Convert to boolean

	port, err := strconv.Atoi(mailPort)
	if err != nil {
		return fmt.Errorf("invalid mail port: %v", err)
	}

	msg := mail.NewMessage()
	msg.SetHeader("From", senderEmail)
	msg.SetHeader("To", recipientEmail)
	msg.SetHeader("Subject", fmt.Sprintf("New Enquiry from %s", userName))
	msg.SetBody("text/plain", fmt.Sprintf("Name: %s\nEmail: %s\nPhone Number: %s\nMessage: %s", userName, userEmail, userPhone, userMessage))

	dialer := mail.NewDialer(mailServer, port, senderEmail, senderPassword)

	// Configure TLS if required
	if useTLS {
		dialer.TLSConfig = &tls.Config{
			ServerName:         mailServer, // Set the server name here
			InsecureSkipVerify: false,      // Set to true only for testing purposes
		}
	}

	if err := dialer.DialAndSend(msg); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
