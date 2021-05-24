package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	Image     string             `bson:"img,omitempty" json:"img"`
	Role      string             `bson:"role" json:"role"`
	Google    bool               `bson:"google,omitempty" json:"google"`
	Status    string             `bson:"status" json:"status"`
	CreatedAt int64              `bson:"created_at" json:"createdOn"`
	UpdateBy  primitive.ObjectID `bson:"updated_by,omitempty" json:"updatedBy"`
	UpdatedAt int64              `bson:"updated_at,omitempty" json:"updatedOn"`
	DeletedAt int64              `bson:"deleted_at,omitempty" json:"deletedOn"`
}

type Users []User
