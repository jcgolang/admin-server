package user_repository

import (
	db "admin-server/database"
	"admin-server/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	ctx        context.Context
	collection *mongo.Collection
}

var repository *Repository

func init() {
	repository = &Repository{
		collection: db.GetManager().GetCollection("users"),
		ctx:        context.Background(),
	}
}

func GetRepository() *Repository {
	return repository
}

func (r *Repository) Create(user *models.User) error {

	if result, err := r.collection.InsertOne(r.ctx, user); err != nil {
		return err
	} else {
		id, _ := result.InsertedID.(primitive.ObjectID)
		user.Id = &id
	}
	return nil
}

func (r *Repository) GetUserById(id string) (*models.User, error) {

	// Convierte el id a un id de mongo
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("error: '%s'es un id invalido", id)
	}

	user := &models.User{}

	// No es un error, si el id no existe, pero retorna nulo
	// TODO: Revisar si se retorna un usuario "vacio" o nulo
	result := r.collection.FindOne(r.ctx, bson.M{"_id": objId})
	if result.Err() != nil {
		return user, nil
	}

	// Realiza la decodificación y si no la puede realizar, returna el error.
	if err := result.Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) GetUsers() (*models.Users, error) {
	cursor, err := r.collection.Find(r.ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)

	users := &models.Users{}
	// TODO: Cambiar a una iteración con for cursor.Next(r.ctx)
	if err := cursor.All(r.ctx, users); err != nil {
		return nil, err
	}
	return users, nil
}
