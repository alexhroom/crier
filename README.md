# file_to_email
Go program that sends a text file as an email.

# Setup
1. [Have Go installed on your machine](https://golang.org/doc/install).
2. Clone the repository.
3. Enter the repository and run `go build .` at the terminal.
4. You should now have an executable called file_to_email in the repository. You are done.

# Usage
In the file credentials.json, fill out the fields with the following:

- `email`: the email address you'd like to send from.
- `server`: the SMTP email server. For example, for gmail this would be `smtp.gmail.com`. Most email services have their SMTP details somewhere on their website.
- `port`: the outgoing port of the SMTP email server. For example, again for gmail this is 587. Again, most services have this on their website.

To then send a content file as an email, from a terminal in the repository type the following:

`file_to_email -file [content file path] -list [mailing list path] [options]`

where `[content file path]` is the path to your content file and `[mailing list path]` is the location of your mailing list file; i.e. a list of email addresses you'd like to send the text file to.

Current supported options are:
- `-cc [address]`: Cc an email address into the email.
- `-mime [mimetype]`: Choose the [MIME type of your file](https://www.sitepoint.com/mime-types-complete-list/). Defauls to `text/plain`. If you're using HTML you should set this to `text/html`, or if using CSS you should set this to `text/css`.

# Development
I'm developing this as a personal project, and would like to build it into a sort of minimalist newsletter client, with better support for formatting & mailing lists. Ideally, users will just be able to run `file_to_email my_newsletter` (or whatever I end up renaming it to. I think Crier might sound catchy) and send a newsletter from the command line.
