package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	IsBlocked bool   `json:"isBlocked,omitempty"`
	Password  string `json:"password,omitempty"`
}

type UserDB struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name,omitempty"`
	Email     string             `bson:"email,omitempty"`
	IsBlocked bool               `bson:"isBlocked,omitempty"`
	Password  string             `bson:"password,omitempty"`
}

func MapUserDB(user UserDB) User {
	return User{
		ID:        user.ID.Hex(),
		Name:      user.Name,
		Email:     user.Email,
		IsBlocked: user.IsBlocked,
	}
}

func CheckPassword(user UserDB, password string) bool {
	currentPassword := []byte(user.Password)
	candidate := []byte(password)

	err := bcrypt.CompareHashAndPassword(currentPassword, candidate)
	if err != nil {
		return false
	} else {
		return true
	}
}
