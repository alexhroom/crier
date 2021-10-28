package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/smtp"
	"path"
	"strconv"
)

func main() {
	// setup flags: path to text file to send, and receiver of file
	filepath := flag.String("filepath", "NA", "The path to the text file that you would like to send as an email.")
	receiver := flag.String("receiver", "NA", "The receiver of the email.")

	flag.Parse()

	if *filepath == "NA" {
		log.Fatal("Please specify flag -filepath.")
	}
	if *receiver == "NA" {
		log.Fatal("Please specify flag -receiver.")
	}

	// read text file and get file name for email subject
	content, read_err := ioutil.ReadFile(*filepath)
	if read_err != nil {
		log.Fatal("Read Error: ", read_err)
	}
	filename := path.Base(*filepath)
	filetext := string(content)

	// authenticate email using credentials in json file
	cred_data, credread_err := ioutil.ReadFile("./credentials.json")
	if credread_err != nil {
		log.Fatal("Credentials read error: ", credread_err)
	}

	type Credentials struct {
		Username string
		Password string
		Server   string
		Port     int
	}

	var credentials Credentials
	auth_err := json.Unmarshal(cred_data, &credentials)

	if auth_err != nil {
		log.Fatal("Auth error: ", auth_err)
	}

	auth := smtp.PlainAuth("", credentials.Username, credentials.Password, credentials.Server)

	// send email
	to := []string{*receiver}
	msg := []byte("To:" + *receiver + "\r\n" +
		"Subject:" + filename + "\r\n" +
		"\r\n" +
		filetext)
	send_err := smtp.SendMail(credentials.Server+":"+strconv.Itoa(credentials.Port), auth, credentials.Username, to, msg)
	if send_err != nil {
		log.Fatal("Send Error: ", send_err)
	}
}
