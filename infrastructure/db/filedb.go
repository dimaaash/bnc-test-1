package db

import (
	"context"
	"log"

	"github.com/dimaaash/bnc-test-1/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBHandler struct {
	MongoClient mongo.Client
	database    *mongo.Database
}

func NewDBHandler(connectString string, dbName string) (DBHandler, error) {
	dbHandler := DBHandler{}
	clientOptions := options.Client().ApplyURI(connectString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return dbHandler, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return dbHandler, err
	}
	dbHandler.MongoClient = *client
	dbHandler.database = client.Database(dbName)
	return dbHandler, nil
}

func (dbHandler DBHandler) FindAllBooks() ([]*domain.Book, error) {
	var results []*domain.Book
	collection := dbHandler.database.Collection("books")
	cur, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var elem domain.Book
		err2 := cur.Decode(&elem)
		if err2 != nil {
			log.Fatal(err2)
		}
		results = append(results, &elem)
	}
	return results, nil
}

func (dbHandler DBHandler) SaveCard(card domain.Card) error {
	collection := dbHandler.database.Collection("cards")
	_, err := collection.InsertOne(context.TODO(), card)
	if err != nil {
		return err
	}
	return nil
}
func (dbHandler DBHandler) SaveAccount(account domain.Account) error {
	collection := dbHandler.database.Collection("accounts")
	_, err := collection.InsertOne(context.TODO(), account)
	if err != nil {
		return err
	}
	return nil
}
