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
	Id        *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string              `bson:"name" json:"name,omitempty"`
	Email     string              `bson:"email" json:"email,omitempty"`
	Password  []byte              `bson:"password" json:"password,omitempty"`
	Image     string              `bson:"img,omitempty" json:"img,omitempty"`
	Role      Role                `bson:"role" json:"role,omitempty"`
	Google    bool                `bson:"google,omitempty" json:"google,omitempty"`
	Status    Status              `bson:"status" json:"status,omitempty"`
	CreatedAt int64               `bson:"created_at" json:"createdOn,omitempty"`
	UpdateBy  *primitive.ObjectID `bson:"updated_by,omitempty" json:"updatedBy,omitempty"`
	UpdatedAt int64               `bson:"updated_at,omitempty" json:"updatedOn,omitempty"`
	DeletedAt int64               `bson:"deleted_at,omitempty" json:"deletedOn,omitempty"`
}

type Users []*User

func NewUser() *User {
	return &User{
		Status:    ACTIVE,
		Role:      USER_ROLE,
		CreatedAt: time.Now().Unix(),
	}
}
