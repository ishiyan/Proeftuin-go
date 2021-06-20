package main

import (
	"fmt"
	"strings"
)

type email struct {
	from, to, subject, body string
}

// EmailBuilder builds e-mail.
type EmailBuilder struct {
	email email
}

// From sets the from address.
func (b *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("email should contain @")
	}
	b.email.from = from
	return b
}

// To sets the to address.
func (b *EmailBuilder) To(to string) *EmailBuilder {
	b.email.to = to
	return b
}

// Subject sets the subject.
func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.email.subject = subject
	return b
}

// Body sets the body.
func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.body = body
	return b
}

func sendMailImpl(email *email) {
	// actually ends the email
	fmt.Println(*email)
}

type build func(*EmailBuilder)

// SendEmail sends a e-mail.
func SendEmail(action build) {
	builder := EmailBuilder{}
	action(&builder)
	sendMailImpl(&builder.email)
}

func main() {
	SendEmail(func(b *EmailBuilder) {
		b.From("foo@bar.com").
			To("bar@baz.com").
			Subject("Meeting").
			Body("Hello, do you want to meet?")
	})
}
