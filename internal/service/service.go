package service

import (
	"bytes"
	"log"
	"text/template"
	"time"

	"github.com/gipsh/stori-challenge/internal/domain"
	"github.com/gipsh/stori-challenge/internal/mailer"
	"github.com/gipsh/stori-challenge/internal/parser"
	"github.com/gipsh/stori-challenge/internal/reader"
	"github.com/gipsh/stori-challenge/internal/repository"
)

type Service struct {
	mailer     mailer.Mailer
	repository repository.Repository
	parser     *parser.Parser
	reader     reader.FileReader
}

func NewService(mailer mailer.Mailer, repo repository.Repository, reader reader.FileReader) Service {
	return Service{mailer: mailer, repository: repo, parser: parser.NewParser(), reader: reader}
}

func (d *Service) ProcessFile(filename string) error {
	txs, err := d.ReadData(filename)
	if err != nil {
		return err
	}

	summary := d.GenerateSummary(txs)

	err = d.SendSummary(summary)
	if err != nil {
		return err
	}

	return nil

}

// ProcessFile processes the file and returns a list of transactions
func (d *Service) ReadData(filename string) ([]domain.Transaction, error) {

	file, err := d.reader.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	txs, err := d.parser.ParseFile(file)
	if err != nil {
		return nil, err
	}

	return txs, nil
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
			log.Println(err)
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
