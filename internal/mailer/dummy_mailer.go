package mailer

import "log"

type DummyMailer struct {
}

func (d *DummyMailer) Send(to, subject, body string) error {
	log.Println("Body:", body)
	return nil
}
