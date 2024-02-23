package mock

import "github.com/gipsh/stori-challenge/internal/domain"

type RepositoryMock struct {
}

func (r *RepositoryMock) CreateTransaction(transaction domain.Transaction) error {
	return nil
}
