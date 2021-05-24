package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	user     = "app"
	password = "123456"
	host     = "localhost"
	port     = 27017
	dbName   = "hospital"
	db       *mongo.Client
	ctx      context.Context
)

func init() {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", user, password, host, port)

	var err error

	db, err = mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err := db.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	defer cancel()
}

func Disconnect() {
	fmt.Println("Cerrando conexión con mongodb...")
	db.Disconnect(ctx)
}

func GetCollection(collection string) *mongo.Collection {
	return db.Database(dbName).Collection(collection)
}
