package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"strconv"
	"syscall"
	"time"

	"golang.org/x/term"
)

// Contains functions to send an Email struct via SMTP.

// Credentials is a struct designed to store credentials.json as an object.
type Credentials struct {
	Email  string
	Server string
	Port   int
}

// Auth reads credentials.json and sets up the user's email authorisation.
func Auth() (Credentials, smtp.Auth) {
	// ask for password
	fmt.Println("Please enter password:")
	password, passErr := term.ReadPassword(int(syscall.Stdin))
	if passErr != nil {
		log.Fatal("Password error: ", passErr)
	}

	// authenticate email using credentials in json file
	credData, credreadErr := ioutil.ReadFile("./credentials.json")
	if credreadErr != nil {
		log.Fatal("Credentials read error: ", credreadErr)
	}

	var credentials Credentials
	authErr := json.Unmarshal(credData, &credentials)

	if authErr != nil {
		log.Fatal("Auth error: ", authErr)
	}
	auth := smtp.PlainAuth("", credentials.Email, string(password), credentials.Server)
	return credentials, auth
}

// SendEmail sets up email headers & payload then sends the email.
func SendEmail(email Email, mimetype string, credentials Credentials, auth smtp.Auth) {
	// create vars for SMTP headers
	now := time.Now()
	to := []string{email.To}
	cc := "Cc:" + email.Cc + "\r\n"
	subject := "Subject:" + email.Subject + "\r\n"
	mime := "MIME-version: 1.0;\nContent-Type: " + mimetype + "; charset=\"UTF-8\";\n\n"
	msg := []byte(
		// SMTP headers are set up here
		"Date:" + now.Format(time.RFC1123Z) + "\r\n" +
			"From:" + credentials.Email + "\r\n" +
			"To:" + email.To + "\r\n" +
			cc +
			subject +
			mime +
			"\r\n" +
			// SMTP payload is email.Message
			email.Message)
	sendErr := smtp.SendMail(credentials.Server+":"+strconv.Itoa(credentials.Port), auth, credentials.Email, to, msg)
	if sendErr != nil {
		log.Fatal("Send error: ", sendErr)
	}
	fmt.Println("Email sent!")
}
