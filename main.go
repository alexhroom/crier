package main

import (
	"flag"
	"log"
	"strings"

	src "github.com/alexhroom/crier/src"
)

func main() {
	// setup flags: path to text file to send, and receiver of file
	file := flag.String("file", "./content", "The path to the text file that you would like to send as an email.")
	mailingList := flag.String("list", "./mailing_list.txt", "The receiver of the email.")
	cc := flag.String("cc", "", "Optionally add an address to cc the file to. This could be your own address.")
	mimetype := flag.String("mime", "text/plain", "Optionally specify the MIME type. This defaults to plaintext, but if you're using HTML then specify text/html.")

	flag.Parse()

	neededFlags := make([]string, 0, 2)
	if *file == "./content" {
		neededFlags = append(neededFlags, "-file (content file)")
	}
	if *mailingList == "./mailing_list.txt" {
		neededFlags = append(neededFlags, "-list (recipient file)")
	}
	if len(neededFlags) > 0 {
		log.Print("WARNING: You haven't set the flags: " + strings.Join(neededFlags, ", ") + ". Did you forget to specify these files?")
	}

	email := src.CreateEmail(*file, *mailingList, *cc)
	credentials, auth := src.Auth()

	src.SendEmail(email, *mimetype, credentials, auth)
}
