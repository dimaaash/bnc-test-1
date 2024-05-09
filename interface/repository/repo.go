package repository

import "github.com/dimaaash/bnc-test-1/domain"

type DBHandler interface {
	FindAllAccounts() ([]*domain.Account, error)
	FindAllCards() ([]*domain.Card, error)

	SaveAccount(account domain.Account) error
	UpdateAccount(account domain.Account) error
	SaveCard(card domain.Card) error
}
