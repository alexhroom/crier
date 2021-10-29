package main

import (
	"flag"
	"log"
	"strings"

	src "github.com/alexhroom/file_to_email/src"
)

func main() {
	// setup flags: path to text file to send, and receiver of file
	file := flag.String("file", "./content", "The path to the text file that you would like to send as an email.")
	mailing_list := flag.String("list", "./mailing_list.txt", "The receiver of the email.")
	cc := flag.String("cc", "", "Optionally add an address to cc the file to. This could be your own address.")
	mimetype := flag.String("mime", "text/plain", "Optionally specify the MIME type. This defaults to plaintext, but if you're using HTML then specify text/html.")

	flag.Parse()

	needed_flags := make([]string, 0, 2)
	if *file == "./content" {
		needed_flags = append(needed_flags, "-file (content file)")
	}
	if *mailing_list == "./mailing_list.txt" {
		needed_flags = append(needed_flags, "-list (recipient file)")
	}
	if len(needed_flags) > 0 {
		log.Print("WARNING: You haven't set the flags: " + strings.Join(needed_flags, ", ") + ". Did you forget to specify these files?")
	}

	email := src.CreateEmail(*file, *mailing_list, *cc)
	credentials, auth := src.Auth()

	src.SendEmail(email, *mimetype, credentials, auth)
}
