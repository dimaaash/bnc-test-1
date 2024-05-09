package usecases

import (
	"log"

	"github.com/dimaaash/bnc-test-1/domain"
)

type CardInteractor struct {
	CardRepository domain.CardRepository
}

func NewCardInteractor(CardRepository domain.CardRepository) CardInteractor {
	return CardInteractor{CardRepository}
}

func (interactor *NewCardInteractor) CreateCard(card domain.Card) error {
	err := interactor.CardRepository.SaveCard(card)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
