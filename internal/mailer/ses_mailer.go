package mailer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

type SESMailer struct {
	client *ses.Client
	from   string
}

func NewSESMailer(cfg aws.Config, from string) *SESMailer {
	return &SESMailer{
		client: ses.NewFromConfig(cfg),
		from:   from,
	}
}

func (s *SESMailer) Send(to, subject, body string) error {

	msg := []byte("To: " + to + "\r\n" +
		"From: " + s.from + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		"\r\n" +
		body)

	input := &ses.SendRawEmailInput{
		RawMessage: &types.RawMessage{
			Data: msg,
		},
	}

	_, err := s.client.SendRawEmail(context.TODO(), input)
	if err != nil {
		return err
	}

	return nil
}
