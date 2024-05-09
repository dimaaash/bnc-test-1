package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/dimaaash/bnc-test-1/domain"

	"github.com/dimaaash/bnc-test-1/usecases"
)

type AccountController struct {
	accountInteractor usecases.AccountInteractor
}

func NewAccountController(accountInteractor usecases.AccountInteractor) *AccountController {
	return &AccountController{accountInteractor}
}

func (controller *AccountController) Add(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var account domain.Account
	err := json.NewDecoder(req.Body).Decode(&account)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid Payload"})
		return
	}
	err2 := controller.accountInteractor.CreateAccount(account)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
}
