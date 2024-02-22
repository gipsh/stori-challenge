package service

import (
	"reflect"
	"testing"

	"github.com/gipsh/stori-challenge/internal/domain"
)

func TestDomain_GenerateSummary(t *testing.T) {

	d := NewService(nil, nil)

	txs := []domain.Transaction{
		{Id: 0, Month: 7, Day: 15, Amount: 60.5, IsDebit: false},
		{Id: 1, Month: 7, Day: 28, Amount: -10.3, IsDebit: true},
		{Id: 2, Month: 8, Day: 2, Amount: -20.46, IsDebit: true},
		{Id: 3, Month: 8, Day: 13, Amount: 10, IsDebit: false},
	}

	summary := d.GenerateSummary(txs)

	expected := domain.Summary{
		TotalBalance:  39.74,
		AverageDebit:  -15.38,
		AverageCredit: 35.25,
		MonthlyTransactions: map[string]int{
			"July":   2,
			"August": 2,
		},
	}

	if !reflect.DeepEqual(summary, expected) {
		t.Errorf("Expected %v, but got %v", expected, summary)
	}

}