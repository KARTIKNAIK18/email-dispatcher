# Email Dispatcher

A Go program that reads recipients from a CSV file and sends a templated
email to each recipient using concurrent workers.

## Requirements

- Go 1.25.3 or newer
- An SMTP server available at `localhost:1025`

### Recommended SMTP server: Mailpit

For local development, use [Mailpit](https://github.com/axllent/mailpit).
Mailpit provides a local SMTP server and a web inbox, so test emails are
captured instead of being delivered to real recipients.

Start Mailpit with Docker:

```bash
docker run -d --name mailpit -p 1025:1025 -p 8025:8025 axllent/mailpit
```

The application sends mail to Mailpit at `localhost:1025`. Open
<http://localhost:8025> to view captured messages.

The default SMTP settings are defined in
[`Pipeline/consumer.go`](./Pipeline/consumer.go). Update the host, port, and
sender address before using a different mail server. For production, use an
authenticated SMTP provider such as Amazon SES, SendGrid, Mailgun, or
Brevo instead of the local Mailpit server. Do not commit SMTP credentials to
the repository.

## Project structure

```text
email-dispatcher/
├── Pipeline/
│   ├── consumer.go   # Sends emails from the queue
│   ├── model.go      # Recipient model
│   ├── producer.go   # Loads recipients from CSV
│   └── template.go   # Renders the email template
├── contact.csv       # Recipient input
├── email.tmpl        # Email subject and body
├── main.go           # Application entry point
├── go.mod
└── README.md
```

## Input CSV

The application expects `contact.csv` in the project root. The first row is
treated as a header, followed by recipient name and email columns:

```csv
name,email
Alice,alice@example.com
Bob,bob@example.com
```

## Email template

The template is loaded from `email.tmpl` and can use the recipient's `Name`
field:

```text
Subject: hello,{{.Name}}

hi {{.Name}}

How are You!!
```

## Run

From the project root, run:

```bash
go run .
```

The program loads the CSV, places recipients on a queue, and starts five
worker slots to process the queue.

## Build

```bash
go build -o email-dispatcher .
```

Run the generated executable from the project root so it can find
`contact.csv` and `email.tmpl`.