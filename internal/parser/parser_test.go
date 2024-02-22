package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTransactionDebit(t *testing.T) {

	parser := NewParser()

	tx, err := parser.ParseTransaction("0,7/15,+60.5")

	assert.Equal(t, nil, err)
	assert.Equal(t, 0, tx.Id)
	assert.Equal(t, 7, tx.Month)
	assert.Equal(t, 15, tx.Day)
	assert.Equal(t, 60.5, tx.Amount)
	assert.Equal(t, false, tx.IsDebit)

}

func TestParseTransactionCredit(t *testing.T) {

	parser := NewParser()

	tx, err := parser.ParseTransaction("0,7/15,-10.5")

	assert.Equal(t, nil, err)
	assert.Equal(t, 0, tx.Id)
	assert.Equal(t, 7, tx.Month)
	assert.Equal(t, 15, tx.Day)
	assert.Equal(t, -10.5, tx.Amount)
	assert.Equal(t, true, tx.IsDebit)

}
