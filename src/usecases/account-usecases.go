package usecases

import (
	"log"

	"github.com/github.com/dimaaash/bnc-test-1/domain"
)

type AccountInteractor struct {
	AccountRepository domain.AccountRepository
}

func NewAccountInteractor(AccountRepository domain.AccountRepository) AccountInteractor {
	return AccountInteractor{AccountRepository}
}

func (interactor *AccountInteractor) CreateAccount(account domain.Account) error {
	err := interactor.AccountRepository.SaveAccount(account)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
