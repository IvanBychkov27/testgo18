package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/go-mail/mail"
)

// Регулярное выражение проверки почты
var rxEmail = regexp.MustCompile(".+@.+\\..+")

type Message struct {
	Email   string
	Content string
	Errors  map[string]string
}

func (msg *Message) Validate() bool {
	msg.Errors = make(map[string]string)

	match := rxEmail.Match([]byte(msg.Email))
	if match == false {
		msg.Errors["Email"] = "Please enter a valid email address"
	}

	if strings.TrimSpace(msg.Content) == "" {
		msg.Errors["Content"] = "Please enter a message"
	}

	return len(msg.Errors) == 0
}

func (msg *Message) Action() error {
	log.Printf("%s %s", msg.Email, msg.Content)
	return nil
}

func (msg *Message) Deliver() error {
	email := mail.NewMessage()
	email.SetHeader("To", "IvBrynsk@rambler.ru")
	email.SetHeader("From", "IvanBychkov@mail.ru")
	email.SetHeader("Reply-To", msg.Email)
	email.SetHeader("Subject", "New message via Contact Form")
	email.SetBody("text/plain", msg.Content)

	username := "username"
	password := "password"

	return mail.NewDialer("smtp.mailtrap.io", 25, username, password).DialAndSend(email)
}
