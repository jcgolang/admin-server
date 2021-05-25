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
	user            = config.Config.Mongo.User
	password        = config.Config.Mongo.Pass
	host            = config.Config.Mongo.Host
	port            = config.Config.Mongo.Port
	dbName          = config.Config.Mongo.DbName
	timeOut         = config.Config.Mongo.TimeOut
	db              *mongo.Client
	usersCollection *mongo.Collection
	ctx             context.Context
	cancel          context.CancelFunc
)

func Connect() {

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", user, password, host, port)

	var err error
	db, err = mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		//panic("Error al crear cliente de mongodb")
		log.Fatal("Error al crear cliente de mongodb.\n", err)
	}

	ctx, cancel = context.WithTimeout(context.TODO(), timeOut*time.Second)
	if err := db.Connect(ctx); err != nil {
		//panic("Error al establecer el contexto global de mongodb")
		log.Fatal(err)
	}
	defer cancel()

	if err := db.Ping(ctx, nil); err != nil {
		//panic(fmt.Sprintf("%s\n%s", "Error al establcer conexión con mongodb", err))
		log.Fatal("Error al establecer conexión con mongodb.\n", err)
	}
	usersCollection = getCollection("users")
}

func Disconnect() {
	fmt.Println("Cerrando conexión con mongodb...")
	if db != nil {
		db.Disconnect(ctx)
	}
}

func getCollection(collection string) *mongo.Collection {
	return db.Database(dbName).Collection(collection)
}

func GetUsersCollection() *mongo.Collection {
	return usersCollection
}
