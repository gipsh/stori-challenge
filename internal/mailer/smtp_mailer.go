package mailer

import (
	"fmt"
	"net/smtp"
)

type SMTPMailer struct {
	Host     string
	Port     string
	Username string
	Password string
	From     string
}

func NewSMTPMailer(host, port, username, password, from string) *SMTPMailer {
	return &SMTPMailer{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		From:     from,
	}
}

func (s *SMTPMailer) Send(to, subject, body string) error {

	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)
	tos := []string{to}
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		"\r\n" +
		body)

	err := smtp.SendMail(fmt.Sprintf("%s:%s", s.Host, s.Port), auth, s.From, tos, msg)

	return err

}
