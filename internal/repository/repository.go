package repository

import (
	"database/sql"

	"github.com/gipsh/stori-challenge/internal/domain"
)

type Repository interface {
	CreateTransaction(transaction domain.Transaction) error
}

type RepositoryImpl struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *RepositoryImpl {
	return &RepositoryImpl{db: db}
}

func (r *RepositoryImpl) CreateTransaction(transaction domain.Transaction) error {
	_, err := r.db.Exec("INSERT INTO transactions (id, month, day, amount, account) VALUES (?, ?, ?, ?, ?)",
		transaction.Id,
		transaction.Month,
		transaction.Day,
		transaction.Amount,
		transaction.Account)
	if err != nil {
		return err
	}

	return nil
}

// create table
// CREATE TABLE transactions (
// 	id INTEGER PRIMARY KEY,
// 	month INTEGER,
// 	day INTEGER,
// 	amount REAL,
// 	account TEXT
// );
