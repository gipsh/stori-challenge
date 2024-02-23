//go:generate mockgen -source=repository.go -destination=mocks/repository_mock.go -package=mock
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
	_, err := r.db.Exec("INSERT INTO transactions ( month, day, amount, account) VALUES ( ?, ?, ?, ?)",
		transaction.Month,
		transaction.Day,
		transaction.Amount,
		transaction.Account)
	if err != nil {
		return err
	}

	return nil
}
