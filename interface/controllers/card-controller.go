package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/dimaaash/bnc-test-1/domain"

	"github.com/dimaaash/bnc-test-1/usecases"
)

type CardController struct {
	cardInteractor usecases.CardInteractor
}

func NewCardController(cardInteractor usecases.CardInteractor) *CardController {
	return &CardController{cardInteractor}
}

func (controller *CardController) Add(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var card domain.Card
	err := json.NewDecoder(req.Body).Decode(&card)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid Payload"})
		return
	}
	err2 := controller.cardInteractor.CreateCard(card)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (controller *CardController) FindAll(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	results, err2 := controller.cardInteractor.FindAll()
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(results)
}
