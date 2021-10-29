// Contains functions that send an Email struct.
package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"strconv"
	"syscall"

	"golang.org/x/term"
)

type Credentials struct {
	Email  string
	Server string
	Port   int
}

// Reads credentials.json and sets up the user's email authorisation.
func Auth() (Credentials, smtp.Auth) {
	// ask for password
	fmt.Println("Please enter password:")
	password, pass_err := term.ReadPassword(int(syscall.Stdin))
	if pass_err != nil {
		log.Fatal("Password error: ", pass_err)
	}

	// authenticate email using credentials in json file
	cred_data, credread_err := ioutil.ReadFile("./credentials.json")
	if credread_err != nil {
		log.Fatal("Credentials read error: ", credread_err)
	}

	var credentials Credentials
	auth_err := json.Unmarshal(cred_data, &credentials)

	if auth_err != nil {
		log.Fatal("Auth error: ", auth_err)
	}
	auth := smtp.PlainAuth("", credentials.Email, string(password), credentials.Server)
	return credentials, auth
}

// Sends an Email
func SendEmail(email Email, credentials Credentials, auth smtp.Auth) {
	// send email
	to := []string{email.To}
	msg := []byte("To:" + email.To + "\r\n" +
		"Subject:" + email.Subject + "\r\n" +
		"\r\n" +
		email.Message)
	send_err := smtp.SendMail(credentials.Server+":"+strconv.Itoa(credentials.Port), auth, credentials.Email, to, msg)
	if send_err != nil {
		log.Fatal("Send error: ", send_err)
	}
	fmt.Println("Email sent!")
}
