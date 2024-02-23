package mailer

import "fmt"

type DummyMailer struct {
}

func (d *DummyMailer) Send(to, subject, body string) error {
	fmt.Println("Body:", body)
	return nil
}
