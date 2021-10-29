package main

import (
	"flag"
	"log"
	"strings"

	src "github.com/alexhroom/file_to_email/src"
)

func main() {
	// setup flags: path to text file to send, and receiver of file
	file := flag.String("file", "NA", "The path to the text file that you would like to send as an email.")
	mailing_list := flag.String("list", "NA", "The receiver of the email.")
	cc := flag.String("cc", "None", "Optionally add an address to cc the file to. This could be your own address.")

	flag.Parse()

	needed_flags := make([]string, 0, 2)
	if *file == "NA" {
		needed_flags = append(needed_flags, "-file (content file)")
	}
	if *mailing_list == "NA" {
		needed_flags = append(needed_flags, "-list (recipient file)")
	}
	if len(needed_flags) > 0 {
		log.Fatal("Mandatory flags " + strings.Join(needed_flags, ", ") + " not supplied.")
	}

	email := src.CreateEmail(*file, *mailing_list, *cc)
	credentials, auth := src.Auth()

	src.SendEmail(email, credentials, auth)
}
