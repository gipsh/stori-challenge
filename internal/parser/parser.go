package parser

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gipsh/stori-challenge/internal/domain"
)

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

// ParseFile process file line by line
func (p *Parser) ParseFile(filename string) ([]domain.Transaction, error) {

	// Open file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	txs := make([]domain.Transaction, 0)

	// Read file line by line and parse transaction
	scanner := bufio.NewScanner(file)
	// skip first line
	scanner.Scan()
	for scanner.Scan() {

		line := scanner.Text()
		tx, err := p.ParseTransaction(line)
		if err != nil {
			return nil, err
		}
		txs = append(txs, tx)
	}

	return txs, nil
}

func (p *Parser) ParseTransaction(line string) (domain.Transaction, error) {

	var tx domain.Transaction

	if _, err := fmt.Sscanf(line, "%d,%d/%d,%f", &tx.Id, &tx.Month, &tx.Day, &tx.Amount); err != nil {
		return tx, err
	}

	if tx.Amount < 0 {
		tx.IsDebit = true
	}

	return tx, nil
}
