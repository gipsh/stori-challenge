package domain

// Summary is a struct that holds the total credit and debit
type Summary struct {
	TotalBalance        float64
	AverageDebit        float64
	AverageCredit       float64
	MonthlyTransactions map[string]int
}

type Transaction struct {
	Month   int
	Day     int
	Id      int
	Amount  float64
	Account string
	IsDebit bool
}
