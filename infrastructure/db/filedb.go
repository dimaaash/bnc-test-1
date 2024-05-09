package db

import (
	"errors"
	"fmt"
	"log"

	"github.com/dimaaash/bnc-test-1/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

type DBHandler struct {
	DB *gorm.DB
	// MongoClient mongo.Client
	// database    *mongo.Database
}

func NewDBHandler(connectString string, dbName string) (DBHandler, error) {

	log.Println("===> DB NewDBHandler <===")

	db, err := gorm.Open(postgres.Open(connectString), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&domain.Card{})
	db.AutoMigrate(&domain.Account{})

	return DBHandler{db}, nil

	// dbHandler := DBHandler{}
	// clientOptions := options.Client().ApplyURI(connectString)
	// client, err := mongo.Connect(context.TODO(), clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return dbHandler, err
	// }
	// err = client.Ping(context.TODO(), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return dbHandler, err
	// }
	// dbHandler.MongoClient = *client
	// dbHandler.database = client.Database(dbName)
	// return dbHandler, nil
}

// FindAllAccounts

func (dbHandler DBHandler) FindAllAccounts() ([]*domain.Account, error) {

	log.Println("===> DB FindAllAccounts <===")

	var accounts []*domain.Account

	// var product domain.Card
	// if result := dbHandler.DB.First(&product); result.Error != nil {
	// 	log.Println("WOOT")
	// }

	if err := dbHandler.DB.Find(&accounts).Error; err != nil {
		// return nil, errors.New("error retriving cards: %s", cards.Error)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Handle not found error
			return nil, errors.New(fmt.Sprintf("error retriving accounts: %s", err))
		} else {
			// Handle other errors
			return nil, errors.New(fmt.Sprintf("error retriving accounts: %s", err))
		}
	}

	// if cards.RowsAffected == 0 {
	// 	return nil, errors.New("No cards found")
	// }
	// if cards.Error != nil {
	// 	return nil, errors.New("error retriving cards: %s", cards.Error)
	// }

	// collection := dbHandler.DB.Collection("books")
	// cur, err := collection.Find(context.TODO(), bson.D{{}})
	// if err != nil {
	// 	return nil, err
	// }
	// for cur.Next(context.TODO()) {
	// 	var elem domain.Book
	// 	err2 := cur.Decode(&elem)
	// 	if err2 != nil {
	// 		log.Fatal(err2)
	// 	}
	// 	results = append(results, &elem)
	// }
	return accounts, nil
}

func (dbHandler DBHandler) FindAllCards() ([]*domain.Card, error) {

	log.Println("===> DB FindAllCards <===")

	var cards []*domain.Card

	// var product domain.Card
	// if result := dbHandler.DB.First(&product); result.Error != nil {
	// 	log.Println("WOOT")
	// }

	if err := dbHandler.DB.Find(&cards).Error; err != nil {
		// return nil, errors.New("error retriving cards: %s", cards.Error)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Handle not found error
			return nil, errors.New(fmt.Sprintf("error retriving cards: %s", err))
		} else {
			// Handle other errors
			return nil, errors.New(fmt.Sprintf("error retriving cards: %s", err))
		}
	}

	// if cards.RowsAffected == 0 {
	// 	return nil, errors.New("No cards found")
	// }
	// if cards.Error != nil {
	// 	return nil, errors.New("error retriving cards: %s", cards.Error)
	// }

	// collection := dbHandler.DB.Collection("books")
	// cur, err := collection.Find(context.TODO(), bson.D{{}})
	// if err != nil {
	// 	return nil, err
	// }
	// for cur.Next(context.TODO()) {
	// 	var elem domain.Book
	// 	err2 := cur.Decode(&elem)
	// 	if err2 != nil {
	// 		log.Fatal(err2)
	// 	}
	// 	results = append(results, &elem)
	// }
	return cards, nil
}

func (dbHandler DBHandler) UpdateAccount(card domain.Account) error {
	log.Println("===> DB UpdateAccount <===")

	// collection := dbHandler.DB.Collection("cards")
	// _, err := collection.InsertOne(context.TODO(), card)
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (dbHandler DBHandler) SaveCard(card domain.Card) error {
	log.Println("===> DB SaveCard <===")

	// collection := dbHandler.DB.Collection("cards")
	// _, err := collection.InsertOne(context.TODO(), card)
	// if err != nil {
	// 	return err
	// }
	return nil
}
func (dbHandler DBHandler) SaveAccount(account domain.Account) error {
	log.Println("===> DB SaveAccount <===")

	// collection := dbHandler.database.Collection("accounts")
	// _, err := collection.InsertOne(context.TODO(), account)
	// if err != nil {
	// 	return err
	// }
	return nil
}
