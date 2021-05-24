package user_repository

import (
	"admin-server/database"
	"admin-server/models"
	"context"
)

var collection = database.GetCollection("users")
var ctx = context.Background()

func Create(user models.User) error {
	if _, err := collection.InsertOne(ctx, user); err != nil {
		return err
	}
	return nil
}
