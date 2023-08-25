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
	ErrUserList     error
	ErrUserFetch    error
	ErrUserGet      error
	ErrUserCreate   error
	ErrUserUpdate   error
	ErrUserDestroy  error
	ErrUserBlock    error
	ErrUserUnblock  error
)

type Users struct {
	collection *mongo.Collection
}

func NewRepositoryUsers(configs *config.Config) *Users {
	ErrUserNotFound = errors.New(configs.Translations.Users.Errors.NotFound)
	ErrUserList = errors.New(configs.Translations.Users.List.Errors)
	ErrUserFetch = errors.New(configs.Translations.Users.Fetch.Errors)
	ErrUserGet = errors.New(configs.Translations.Users.Load.Errors)
	ErrUserCreate = errors.New(configs.Translations.Users.Create.Errors)
	ErrUserUpdate = errors.New(configs.Translations.Users.Update.Errors)
	ErrUserDestroy = errors.New(configs.Translations.Users.Destroy.Errors)
	ErrUserBlock = errors.New(configs.Translations.Users.Block.Errors)
	ErrUserUnblock = errors.New(configs.Translations.Users.Unblock.Errors)

	return &Users{
		collection: configs.Database.Collection("users"),
	}
}

func (rep *Users) ListUsers() ([]types.User, error) {
	items := []types.User{}

	result, err := rep.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, ErrUserList
	}

	defer result.Close(context.Background())

	for result.Next(context.Background()) {
		var item types.UserDB
		if err := result.Decode(&item); err != nil {
			return nil, ErrUserFetch
		}
		items = append(items, types.MapUserDB(item))
	}
	return items, nil
}

func (rep *Users) CreateUser(input types.UserDB) (types.User, error) {
	result, err := rep.collection.InsertOne(context.Background(), input)
	if err != nil {
		return types.User{}, ErrUserCreate
	}
	input.ID = result.InsertedID.(primitive.ObjectID)
	return types.MapUserDB(input), nil
}

func (rep *Users) GetUserByID(id primitive.ObjectID) (types.User, error) {
	var result types.UserDB
	err := rep.collection.
		FindOne(context.Background(), bson.M{"_id": id}).
		Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return types.User{}, ErrUserNotFound
		}
		return types.User{}, ErrUserGet
	}
	return types.MapUserDB(result), nil
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

func (rep *Users) UpdateUser(id primitive.ObjectID, input types.UserDB) error {
	object := bson.M{}
	if input.Name != "" {
		object["name"] = input.Name
	}
	if input.TokenResetPassword != "" {
		object["tokenResetPassword"] = input.TokenResetPassword
	}
	if input.Password != "" {
		object["password"] = input.Password
	}
	result, err := rep.collection.
		UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": object})
	if err != nil {
		return ErrUserUpdate
	}
	if result.MatchedCount == 0 {
		return ErrUserNotFound
	}
	return nil
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

func (rep *Users) BlockOrUnlockUser(id primitive.ObjectID, isBlocked bool) error {
	object := bson.M{}
	object["isBlocked"] = isBlocked

	result, err := rep.collection.
		UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": object})
	if err != nil {
		if isBlocked {
			return ErrUserBlock
		} else {
			return ErrUserUnblock
		}
	}
	if result.MatchedCount == 0 {
		return ErrUserNotFound
	}
	return nil
}
