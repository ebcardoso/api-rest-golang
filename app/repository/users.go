package repository

import (
	"context"
	"errors"

	"github.com/ebcardoso/api-rest-golang/app/types"
	"github.com/ebcardoso/api-rest-golang/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrUserNotFound error
	ErrUserGet      error
	ErrUserCreate   error
	ErrUserDestroy  error
)

type Users struct {
	collection *mongo.Collection
}

func NewRepositoryUsers(configs *config.Config) *Users {
	ErrUserNotFound = errors.New(configs.Translations.Users.Errors.NotFound)
	ErrUserGet = errors.New(configs.Translations.Users.Load.Errors)
	ErrUserCreate = errors.New(configs.Translations.Users.Create.Errors)
	ErrUserDestroy = errors.New(configs.Translations.Users.Destroy.Errors)

	return &Users{
		collection: configs.Database.Collection("users"),
	}
}

func (rep *Users) CreateUser(input types.UserDB) (types.User, error) {
	result, err := rep.collection.InsertOne(context.Background(), input)
	if err != nil {
		return types.User{}, ErrUserCreate
	}
	input.ID = result.InsertedID.(primitive.ObjectID)
	return types.MapUserDB(input), nil
}

func (rep *Users) GetUserByEmail(email string) (types.UserDB, error) {
	var result types.UserDB
	err := rep.collection.
		FindOne(context.Background(), bson.M{"email": email}).
		Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return types.UserDB{}, ErrUserNotFound
		}
		return types.UserDB{}, ErrUserGet
	}
	return result, nil
}

func (rep *Users) DestroyUser(id primitive.ObjectID) error {
	result, err := rep.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return ErrUserDestroy
	}
	if result.DeletedCount == 0 {
		return ErrUserNotFound
	}
	return nil
}
