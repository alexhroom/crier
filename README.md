# file_to_email
Go program that sends a text file as an email.

# Setup
1. [Have Go installed on your machine](https://golang.org/doc/install).
2. Clone the repository.
3. Enter the repository and run `go build .` at the terminal.
4. You should now have an executable called file_to_email in the repository. You are done.

# Usage
Currently, file_to_email uses insecure authorisation to send emails. This means it requires you to write your password in plain text in a file. I don't need to tell you why this is a bad idea, but I wouldn't suggest using file_to_email in its current state for anything important.

In the file credentials.json, fill out the fields with the following:

- `username`: the email address you'd like to send from.
- `password`: the password for this email address. (will be replaced when authorisation is better)
- `server`: the SMTP email server. For example, for gmail this would be `smtp.gmail.com`. Most email services have their SMTP details somewhere on their website.
- `port`: the outgoing port of the SMTP email server. For example, again for gmail this is 587. Again, most services have this on their website.

To then send a file as an email, from a terminal in the repository type the following:

`file_to_email -filepath [path] -receiver [email]`

where [path] is the path to your text file and [email] is the email you'd like to send to.

# Development
I'm developing this as a personal project, and would like to build it into a sort of minimalist newsletter client, with better support for formatting & mailing lists. Ideally, users will just be able to run `file_to_email my_newsletter` (or whatever I end up renaming it to. I think Crier might sound catchy) and send a newsletter from the command line.
