package service

import (
	"bytes"
	"fmt"
	"text/template"
	"time"

	"github.com/gipsh/stori-challenge/internal/domain"
	"github.com/gipsh/stori-challenge/internal/mailer"
	"github.com/gipsh/stori-challenge/internal/repository"
)

type Service struct {
	mailer     mailer.Mailer
	repository repository.Repository
}

func NewService(mailer mailer.Mailer, repo repository.Repository) Service {
	return Service{mailer: mailer, repository: repo}
}

// GenerateSummary generates a summary of the transactions
func (d *Service) GenerateSummary(txs []domain.Transaction) domain.Summary {

	var summary domain.Summary
	summary.MonthlyTransactions = make(map[int]int)
	debits := 0
	credits := 0

	for _, tx := range txs {
		summary.TotalBalance += tx.Amount
		if tx.IsDebit {
			summary.AverageDebit += tx.Amount
			debits++
		} else {
			summary.AverageCredit += tx.Amount
			credits++
		}
		summary.MonthlyTransactions[tx.Month]++

		err := d.repository.CreateTransaction(tx)
		if err != nil {
			fmt.Println(err)
		}

	}

	summary.AverageDebit = summary.AverageDebit / float64(debits)
	summary.AverageCredit = summary.AverageCredit / float64(credits)

	return summary
}

// SendSummary sends the summary to the user
func (d *Service) SendSummary(summary domain.Summary) error {

	// apply the template
	funcs := template.FuncMap{
		"ToMonth": func(month int) string {
			return time.Month(month).String()
		},
	}

	tmpl, err := template.New("summary.tmpl").Funcs(funcs).ParseFiles("internal/domain/template/summary.tmpl")
	if err != nil {
		return err
	}

	var output bytes.Buffer
	err = tmpl.Execute(&output, summary)
	if err != nil {
		return err
	}

	// send the email
	return d.mailer.Send("test@test.com", "Summary", output.String())

}
