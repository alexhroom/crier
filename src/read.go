package src

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Contains functions that set up the email struct.

// Email is a struct containing the content of an email
type Email struct {
	To      string
	Cc      string
	Subject string
	Message string
}

// CreateEmail takes the arguments needed to create an email,
// and turns them into an Email struct
func CreateEmail(ContentFile string, RecipientsFile string, Cc string) Email {
	var email Email

	// get recipients list
	recipients, readRecipientsErr := os.Open(RecipientsFile)
	content, readContentErr := os.Open(ContentFile)
	if readRecipientsErr != nil {
		fmt.Println("Recipient file opening error: ", readRecipientsErr)
	}
	if readContentErr != nil {
		fmt.Println("Content file opening error: ", readContentErr)
	}
	defer recipients.Close()
	defer content.Close()

	// create empty list to hold recipients
	recipientsSlice := make([]string, 0, 999)

	// add mailing list to slice
	recipientsScanner := bufio.NewScanner(recipients)
	for recipientsScanner.Scan() {
		recipientsSlice = append(recipientsSlice, recipients_scanner.Text())
	}

	// convert mailing list to string and add to `To` field
	// also add Cc if available
	email.To = strings.Join(recipientsSlice, ", ")
	email.Cc = Cc

	// create empty slice to hold content
	contentSlice := make([]string, 0, 9999)

	// read text file and get message and subject
	contentScanner := bufio.NewScanner(content)
	for contentScanner.Scan() {
		contentSlice = append(contentSlice, content_scanner.Text())
	}
	// set first line of email as subject
	email.Subject = striphtml(contentSlice[0])
	email.Message = strings.Join(contentSlice, "\r\n")

	return email
}

// Strips HTML from a string.
func striphtml(in string) string {
	// regex to match html tag
	const pattern = `(<\/?[a-zA-A]+?[^>]*\/?>)*`
	r := regexp.MustCompile(pattern)
	groups := r.FindAllString(in, -1)
	for _, group := range groups {
		if strings.TrimSpace(group) != "" {
			in = strings.ReplaceAll(in, group, "")
		}
	}
	return in
}
