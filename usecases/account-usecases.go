package usecases

import (
	"log"

	"github.com/dimaaash/bnc-test-1/domain"
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

func (interactor *AccountInteractor) FindAllAccounts() ([]*domain.Account, error) {
	results, err := interactor.AccountRepository.FindAllAccounts()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	log.Println("results: %+v", results)

	return results, nil
}

func (interactor *AccountInteractor) UpdateAccount(account domain.Account) error {
	err := interactor.AccountRepository.UpdateAccount(account)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
