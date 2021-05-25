package user_repository

import (
	db "admin-server/database"
	"admin-server/models"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ctx = context.Background()

func Create(user *models.User) error {

	var collection = db.GetUsersCollection()

	if id, err := collection.InsertOne(ctx, user); err != nil {
		return err
	} else {
		user.Id = id.InsertedID.(primitive.ObjectID)
	}
	return nil
}
