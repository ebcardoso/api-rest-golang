package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserDB struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
}

func MapUserDB(user UserDB) User {
	return User{
		ID:    user.ID.Hex(),
		Name:  user.Name,
		Email: user.Email,
	}
}
