package repository

import "github.com/dimaaash/bnc-test-1/domain"

type DBHandler interface {
	FindAllAccount() ([]*domain.Book, error)
	SaveAccount(account domain.Account) error
	SaveCard(card domain.Card) error
}
