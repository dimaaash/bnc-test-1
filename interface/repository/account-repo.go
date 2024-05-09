package repository

import (
	"github.com/dimaaash/bnc-test-1/domain"
)

type AccountRepo struct {
	handler DBHandler
}

func NewAccountRepo(handler DBHandler) AccountRepo {
	return AccountRepo{handler}
}

func (repo AccountRepo) SaveAccount(account domain.Account) error {
	err := repo.handler.SaveAccount(account)
	if err != nil {
		return err
	}
	return nil
}