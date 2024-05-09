package repository

import (
	"github.com/dimaaash/bnc-test-1/domain"
)

type CardRepo struct {
	handler DBHandler
}

func NewCardRepo(handler DBHandler) CardRepo {
	return CardRepo{handler}
}

func (repo CardRepo) SaveCard(card domain.Card) error {
	err := repo.handler.SaveCard(card)
	if err != nil {
		return err
	}
	return nil
}

func (repo CardRepo) FindAll() ([]*domain.Card, error) {
	results, err := repo.handler.FindAllCards()
	if err != nil {
		return results, err
	}
	return results, nil
}
