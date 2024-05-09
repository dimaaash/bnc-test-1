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

func (interactor *CardInteractor) CreateCard(card domain.Card) error {
	err := interactor.CardRepository.SaveCard(card)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

// FindAllCards

func (interactor *CardInteractor) FindAllCards() ([]*domain.Card, error) {
	results, err := interactor.CardRepository.FindAllCards()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	log.Println("results: %+v", results)

	return results, nil
}
