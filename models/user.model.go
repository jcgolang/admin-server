package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role string
type Status string

const (
	USER_ROLE  Role   = "USER_ROLE"
	ADMIN_ROLE Role   = "ADMIN_ROLE"
	ACTIVE     Status = "A"
	INACTIVE   Status = "I"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	Password  []byte             `bson:"password" json:"password"`
	Image     string             `bson:"img,omitempty" json:"img"`
	Role      Role               `bson:"role" json:"role"`
	Google    bool               `bson:"google,omitempty" json:"google"`
	Status    Status             `bson:"status" json:"status"`
	CreatedAt int64              `bson:"created_at" json:"createdOn"`
	UpdateBy  primitive.ObjectID `bson:"updated_by,omitempty" json:"updatedBy"`
	UpdatedAt int64              `bson:"updated_at,omitempty" json:"updatedOn"`
	DeletedAt int64              `bson:"deleted_at,omitempty" json:"deletedOn"`
}

type Users []User

func NewUser() *User {
	return &User{
		Status:    ACTIVE,
		Role:      USER_ROLE,
		CreatedAt: time.Now().Unix(),
	}
}
