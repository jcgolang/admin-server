package database

import (
	"admin-server/config"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	user     = config.Config.Mongo.User
	password = config.Config.Mongo.Pass
	host     = config.Config.Mongo.Host
	port     = config.Config.Mongo.Port
	dbName   = config.Config.Mongo.DbName
	timeOut  = config.Config.Mongo.TimeOut
	db       *mongo.Client
	ctx      context.Context
	cancel   context.CancelFunc
)

func init() {

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", user, password, host, port)

	var err error

	db, err = mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Error al crear cliente de mongodb.\n", err)
	}

	ctx, cancel = context.WithTimeout(context.TODO(), timeOut*time.Second)
	if err := db.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(ctx, nil); err != nil {
		log.Fatal("Error al establecer conexión con mongodb.\n", err)
	}
}

func Disconnect() {
	fmt.Println("Cerrando conexión con mongodb...")
	cancel()
	db.Disconnect(ctx)
}

func GetCollection(collection string) *mongo.Collection {
	return db.Database(dbName).Collection(collection)
}
