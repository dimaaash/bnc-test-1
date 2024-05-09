package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dimaaash/bnc-test-1/infrastructure/db"

	"github.com/dimaaash/bnc-test-1/usecases"

	"github.com/dimaaash/bnc-test-1/interface/controllers"

	"github.com/dimaaash/bnc-test-1/infrastructure/router"
	"github.com/dimaaash/bnc-test-1/interface/repository"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
	dbHandler  db.DBHandler
)

func getCardController() controllers.CardController {
	cardRepo := repository.NewCardRepo(dbHandler)
	cardInteractor := usecases.NewCardInteractor(cardRepo)
	cardController := controllers.NewCardController(cardInteractor)
	return *cardController
}

func getAccountController() controllers.AccountController {
	accountRepo := repository.NewAccountRepo(dbHandler)
	accountInteractor := usecases.NewAccountInteractor(accountRepo)
	accountController := controllers.NewAccountController(accountInteractor)
	return *accountController
}

func main() {

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "App is up and running..")
	})
	var err error
	dbHandler, err = db.NewDBHandler("mongodb://localhost:27017", "bookstore")
	if err != nil {
		log.Println("Unable to connect to the DataBase")
		return
	}
	cardController := getCardController()
	accountController := getAccountController()
	httpRouter.POST("/card/add", cardController.Add)
	httpRouter.GET("/card", accountController.FindAll)
	httpRouter.POST("/account/add", accountController.Add)
	httpRouter.SERVE(":8000")

	fmt.Println("hello")
}
