package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Sender's email address and password
	// we need to configure app password on google account setting to get password email sender
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	senderEmail := os.Getenv("SENDER_EMAIL")
	senderPassword := os.Getenv("SENDER_EMAIL_PASS")

	// Recipient's email address
	recipientEmail := "exampleto@gmail.com"

	// SMTP server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := "587" // Use the appropriate SMTP port for your server

	// Message to be sent
	message := []byte("To: " + recipientEmail + "\r\n" +
		"Subject: Test Email\r\n" +
		"\r\n" +
		"This is the body of the email.")

	// Authenticate with the SMTP server
	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)
	fmt.Println(senderEmail, "\n", senderPassword)
	// Send the email
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, []string{recipientEmail}, message)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Email sent successfully!")
}
